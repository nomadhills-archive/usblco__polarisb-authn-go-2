package httpEndpoints

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func getAuthorizationTokenFromHeader(c *gin.Context) (string, error) {
	// get the auth token from the request
	const BEARER_SCHEMA = "Bearer"
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization was null")
	}

	// Get the token string from the header
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]
	return tokenString, nil
}

func getRefreshTokenFromCookies(c *gin.Context) (string, error) {
	// get the auth token from the request
	const REFRESH_TOKEN_COOKIE_NAME = "rft"
	refreshToken, err := c.Cookie(REFRESH_TOKEN_COOKIE_NAME)
	if err != nil {
		return "", errors.New("Refresh token was null")
	}

	return refreshToken, nil
}
