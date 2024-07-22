package middleware

import (
	"errors"
	"net/http"
	"smart-kost-backend/constant"
	"smart-kost-backend/errs"
	"smart-kost-backend/util/response"
	"smart-kost-backend/util/tokenprovider"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateAuth(tokenChecker tokenprovider.JWTTokenProvider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		tokenStr, err := extractToken(authHeader)
		if errors.Is(err, errs.InvalidBearerFormat) {
			response.Error(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		claims, err := tokenChecker.ValidateToken(tokenStr)
		if errors.Is(err, errs.InvalidToken) || errors.Is(err, errs.InvalidIssuer) {
			response.Error(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		if err != nil {
			response.UnknownError(ctx, err)
			return
		}

		ctx.Set(constant.ContextKeyUser, claims.UserClaims)
		ctx.Next()
	}
}

func extractToken(authHeader string) (string, error) {
	splits := strings.Split(authHeader, " ")
	println(authHeader)
	println(splits)

	if len(splits) != 2 || splits[0] != "Bearer" {
		return "", errs.InvalidBearerFormat
	}

	return splits[1], nil
}
