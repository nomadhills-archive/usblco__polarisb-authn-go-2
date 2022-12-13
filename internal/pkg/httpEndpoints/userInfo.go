package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
)

func (endpoints *Endpoints) UserInfoEndpoint(c *gin.Context) {
	user, state, _ := endpoints.Actions.UserGetAuthenticatedUserFromContext(c)
	if state != pkg.UserAuthenticated {
		c.JSON(401, &contracts.UserInfoEndpointReturnDto{
			Err: "User not authenticated",
		})
	}

	// get if user email is verified
	state = endpoints.Actions.UserIsEmailVerified(user)
	if state == pkg.UserEmailVerified {
		c.JSON(200, &contracts.UserInfoEndpointReturnDto{
			Email:         user.Email,
			EmailVerified: true,
		})
	} else {
		c.JSON(200, &contracts.UserInfoEndpointReturnDto{
			Email:         user.Email,
			EmailVerified: false,
		})
	}

}
