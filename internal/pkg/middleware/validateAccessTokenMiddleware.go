package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/usblco/polarisb-authn-go/internal/pkg/actions"
	"github.com/usblco/polarisb-authn-go/pkg"
	"net/http"
)

func (middleware *Middleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		DirectNativeAuthCheck(c, middleware.actions)
	}
}

func (middleware *Middleware) ValidateAccessTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		DirectNativeAuthCheck(c, middleware.actions)
	}
}

func DirectNativeAuthCheck(c *gin.Context, actions *actions.Actions) {
	// get the auth token from the request
	const BEARER_SCHEMA = "Bearer"
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		fmt.Println("Authorization was null")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Get the token string from the header
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	token, state, _ := actions.AccessTokenValidate(tokenString)
	if token == nil {
		if state == pkg.JwtTokenExpired {
			// abort with unauthorized and send error message
			//fmt.Printf("Error validating token: %s \n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Token expired"})
			return
		}
		if state == pkg.InvalidJwtToken {
			// abort with unauthorized and send error message
			//fmt.Printf("Error validating token: %s \n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Invalid token"})
			return
		}
		// abort with unauthorized and send error message
		// write to console
		//fmt.Printf("Error validating token: %s \n", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Invalid token, please login again"})
		return
	}

	if token != nil {
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			userId := claims["id"].(string)
			c.Set("current_user", userId)
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
