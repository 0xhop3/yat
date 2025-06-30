package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/0xhop3/yat/backend/internal/config"
	"github.com/0xhop3/yat/backend/internal/models"
	"github.com/0xhop3/yat/backend/internal/services"
	"github.com/0xhop3/yat/backend/internal/utils"
	"github.com/MicahParks/keyfunc/v2"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationMiddleware struct {
	config      *config.Config
	userService *services.UserService
	jwks        *keyfunc.JWKS
}

type CustomClaims struct {
	Sub      string `json:"sub"`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAuthenticationMiddleware(config *config.Config, userService *services.UserService) *AuthenticationMiddleware {
	// Create JWKS client
	jwksURL := fmt.Sprintf("https://%s/.well-known/jwks.json", config.Auth0Domain)
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshTimeout:  time.Second * 10,
	})
	if err != nil {
		// Log error but don't fail immediately - will retry on first request
		fmt.Printf("Warning: Failed to initialize JWKS: %v\n", err)
	}

	return &AuthenticationMiddleware{
		config:      config,
		userService: userService,
		jwks:        jwks,
	}
}

func (a *AuthenticationMiddleware) ValidateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := a.extractToken(ctx)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid authorization header", err)
			ctx.Abort()
			return
		}

		claims, err := a.validateToken(token)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid token", err)
			ctx.Abort()
			return
		}

		user, err := a.getOrCreateUser(claims)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user", err)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Set("user_id", user.ID)
		ctx.Set("auth0_id", user.Auth0ID)

		ctx.Next()
	}
}

func (a *AuthenticationMiddleware) extractToken(ctx *gin.Context) (string, error) {
	authenticationHeader := ctx.GetHeader("Authorization")
	if authenticationHeader == "" {
		return "", fmt.Errorf("Authorization header required")
	}

	bearerToken := strings.Split(authenticationHeader, "")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return "", fmt.Errorf("Invalid authorization header format")
	}

	return bearerToken[1], nil
}

func (a *AuthenticationMiddleware) validateToken(tokenString string) (*CustomClaims, error) {
	// Initialize JWKS if not done yet
	if a.jwks == nil {
		jwksURL := fmt.Sprintf("https://%s/.well-known/jwks.json", a.config.Auth0Domain)
		jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{
			RefreshInterval: time.Hour,
			RefreshTimeout:  time.Second * 10,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get JWKS: %w", err)
		}
		a.jwks = jwks
	}

	// Parse and validate token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, a.jwks.Keyfunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Verify token is valid
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	// Extract claims
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract claims")
	}

	// Verify issuer
	expectedIssuer := fmt.Sprintf("https://%s/", a.config.Auth0Domain)
	if claims.Issuer != expectedIssuer {
		return nil, fmt.Errorf("invalid issuer")
	}

	return claims, nil
}

func (a *AuthenticationMiddleware) getOrCreateUser(claims *CustomClaims) (*models.User, error) {
	// Try to find existing user by Auth0 ID
	user, err := a.userService.GetByAuth0ID(claims.Sub)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	createRequest := &models.CreateUserRequest{
		Auth0ID:  claims.Sub,
		Username: "N13yx",
		Name:     claims.Name,
	}

	return a.userService.CreateUser(createRequest)
}
