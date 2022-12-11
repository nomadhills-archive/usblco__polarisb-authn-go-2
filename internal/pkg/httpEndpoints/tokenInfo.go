package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/contracts"
)

func (endpoints *Endpoints) TokenInfoEndpoint(c *gin.Context) {
	// Get user from context
	user, state, _ := endpoints.Actions.UserGetAuthenticatedUserFromContext(c)
	if state != pkg.UserAuthenticated {
		c.JSON(401, &contracts.TokenInfoEndpointReturnDto{
			Err: "User not authenticated",
		})
		return
	}

	// Get the token from the header
	tokenString, err := getAuthorizationTokenFromHeader(c)
	if err != nil {
		c.JSON(401, &contracts.TokenInfoEndpointReturnDto{
			Err: "Access token not found",
		})
		return
	}

	// Get the refresh token from the cookies
	/*refreshToken, err := getRefreshTokenFromCookies(c)
	if err != nil {
		c.JSON(401, &contracts.TokenInfoEndpointReturnDto{
			Err: "Refresh token not found",
		})
		return
	}*/

	// Get the token info
	tokenInfo, state, err := endpoints.Actions.AccessTokenGetInfo(tokenString)
	if state != pkg.Ok {
		c.JSON(401, &contracts.TokenInfoEndpointReturnDto{
			Err: "Access token not valid",
		})
		return
	}

	// return json response
	c.JSON(200, &contracts.TokenInfoEndpointReturnDto{
		CurrentUserId:   user.Id.String(),
		AccessTokenInfo: tokenInfo,
	})
}
