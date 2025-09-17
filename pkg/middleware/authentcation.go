package middleware

import (
	"strings"

	"github.com/Farhan1033/resep-masakan-monolith.git/pkg/errs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ContextKey string

const (
	UserIDKey    ContextKey = "user_id"
	UserEmailKey ContextKey = "user_email"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			errsFormat := errs.NewUnauthorized("Authorization header is required")
			ctx.JSON(errsFormat.StatusCode(), errsFormat)
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			errsFormat := errs.NewUnauthorized("Invalid authorization header format")
			ctx.JSON(errsFormat.StatusCode(), errsFormat)
			ctx.Abort()
			return
		}

		tokenString := tokenParts[1]

		_, claims, err := ParseToken(tokenString)
		if err != nil {
			errsFormat := errs.NewUnauthorized(err.Error())
			ctx.JSON(errsFormat.StatusCode(), errsFormat)
			ctx.Abort()
			return
		}

		ctx.Set(string(UserIDKey), claims.ID)
		ctx.Set(string(UserEmailKey), claims.Email)
	}

}

func GetUserIDFromContext(ctx *gin.Context) (uuid.UUID, errs.ErrMessage) {
	userIDStr, exists := ctx.Get(string(UserIDKey))
	if !exists {
		return uuid.Nil, errs.NewNotFound("User ID not found in context")
	}

	userIDString, ok := userIDStr.(string)
	if !ok {
		return uuid.Nil, errs.NewBadRequest("Invalid user ID format in context")
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, errs.NewInternalServerError(err.Error())
	}

	return userID, nil
}
