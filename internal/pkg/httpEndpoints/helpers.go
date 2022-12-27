package httpEndpoints

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go-2/internal/pkg/actions"
	"github.com/usblco/polarisb-authn-go-2/pkg"
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

func checkIfAuthorizedByRole(role string, a *actions.Actions, c *gin.Context) bool {
	user, state, _ := a.UserGetAuthenticatedUserFromContext(c)
	if state != pkg.UserAuthenticated {
		return false
	}

	// check if user has role
	isAuthorized, _, _ := a.CheckIfAuthorizedByRole(role, user)
	return isAuthorized
}
