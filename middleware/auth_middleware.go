package middleware

import (
	"fmt"
	"strings"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/config"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/usecase"
	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const schema = "Bearer"

func validateToken(encodedToken string) (*util.IdTokenClaims, error) {
	claims := &util.IdTokenClaims{}
	token, err := jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}

func AuthorizeJWT(u usecase.UserUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		s := strings.Split(authHeader, fmt.Sprintf("%v ", schema))
		
		authError := util.UnauthorizedError()
		
		if len(s) < 2 {
			ctx.AbortWithStatusJSON(authError.StatusCode, authError)
			return
		}

		decodedToken := s[1]
		claims, err := validateToken(decodedToken)
		if err != nil {
			ctx.AbortWithStatusJSON(authError.StatusCode, authError)
			return
		}

		userId := claims.UserId
		isAdmin := claims.IsAdmin

		ctx.Set("user_id", userId)
		ctx.Set("is_admin", isAdmin)
		ctx.Next()
	}
}
