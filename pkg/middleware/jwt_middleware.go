package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/Farhan1033/resep-masakan-monolith.git/infra/config"
	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const TokenExpiry = 24 * time.Hour

type CustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func CreateToken(id uuid.UUID, email string) (string, errs.ErrMessage) {
	secret := config.GetKey("JWT_SECRET")
	if secret == "" {
		return "", errs.NewInternalServerError("JWT secret not configured")
	}

	claims := jwt.MapClaims{
		"id":    id.String(),
		"email": email,
		"exp":   time.Now().Add(TokenExpiry).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errs.NewInternalServerError(fmt.Sprintf("Failed to sign jwt token: %v", err))
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *CustomClaims, errs.ErrMessage) {
	if tokenString == "" {
		return nil, nil, errs.NewUnauthorized("Token is required")
	}

	secret := config.GetKey("JWT_SECRET")
	if secret == "" {
		return nil, nil, errs.NewInternalServerError("JWT secret not configured")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, nil, errs.NewUnauthorized("Token has expired")
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, nil, errs.NewUnauthorized("Token not valid yet")
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, nil, errs.NewBadRequest("Malformed token")
		case errors.Is(err, jwt.ErrSignatureInvalid):
			return nil, nil, errs.NewUnauthorized("Invalid token signature")
		default:
			return nil, nil, errs.NewUnauthorized(fmt.Sprintf("Token validation failed: %v", err))
		}
	}

	if !token.Valid {
		return nil, nil, errs.NewUnauthorized("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errs.NewBadRequest("Invalid token claims format")
	}

	userID, ok := claims["id"].(string)
	if !ok || userID == "" {
		return nil, nil, errs.NewBadRequest("Invalid user ID in token")
	}

	if _, err := uuid.Parse(userID); err != nil {
		return nil, nil, errs.NewBadRequest("Invalid user ID format in token")
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return nil, nil, errs.NewBadRequest("Invalid email in token")
	}

	parsedClaims := &CustomClaims{
		ID:    userID,
		Email: email,
	}

	return token, parsedClaims, nil
}
