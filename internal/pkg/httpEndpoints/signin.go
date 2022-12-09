package httpEndpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/contracts"
)

func (endpoints *Endpoints) SigninEndpoint(c *gin.Context) {
	// Bind the request to the struct
	request := &contracts.SigninEndpointReceiveUserInfoDto{}
	c.BindJSON(request)

	// Validate the request first.
	if request.Username == "" || request.Password == "" {
		// Return json response
		c.JSON(400, contracts.SigninEndpointReturnDto{
			Err: "username and password are required",
		})
		return
	}

	// Signin the user
	user, state, err := endpoints.Actions.UserSigninByEmail(request.Username, request.Password)
	if err != nil {
		// Return json response
		c.JSON(500, contracts.SigninEndpointReturnDto{
			Err: "Something went wrong",
		})
	}
	if state == pkg.InvalidPasswd {
		// Return json response
		c.JSON(401, contracts.SigninEndpointReturnDto{
			Err: "Please check your email and password",
		})
		return
	}
	if state == pkg.UserNotFound {
		// Return json response
		c.JSON(401, contracts.SigninEndpointReturnDto{
			Err: "Please check your email and password",
		})
		return
	}
	if state != pkg.UserSigninSuccess {
		// Return json response
		c.JSON(500, contracts.SigninEndpointReturnDto{
			Err: "Something went wrong",
		})
		return
	}

	// Generate a JWT access token
	accessToken, state, err := endpoints.Actions.AccessTokenCreate(user)
	if err != nil {
		// Return json response
		c.JSON(500, contracts.SigninEndpointReturnDto{
			Err: "Something went wrong",
		})
		return
	}
	if state != pkg.AccessTokenCreated {
		// Return json response
		c.JSON(500, contracts.SigninEndpointReturnDto{
			Err: "Something went wrong",
		})
		return
	}

	// Generate a JWT refresh token
	// TODO: Implement refresh token

	// Return json response
	c.JSON(200, contracts.SigninEndpointReturnDto{
		Message:     fmt.Sprintf("Welcome %s", user.Email),
		AccessToken: accessToken,
	})
}
