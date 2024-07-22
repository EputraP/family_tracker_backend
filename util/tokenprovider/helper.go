package tokenprovider

import (
	"os"
	"smart-kost-backend/constant"
	"strconv"
)

func GetProvider() JWTTokenProvider {
	issuer := constant.Issuer
	secret := os.Getenv(constant.EnvKeyJWTSecret)
	refreshTokenDurationString := os.Getenv(constant.EnvKeyRefreshTokenDuration)
	accessTokenDurationString := os.Getenv(constant.EnvKeyAccessTokenDuration)

	refreshTokenDuration, _ := strconv.Atoi(refreshTokenDurationString)
	accessTokenDuration, _ := strconv.Atoi(accessTokenDurationString)

	jwtProvider := NewJWT(issuer, secret, refreshTokenDuration, accessTokenDuration)
	return jwtProvider
}
