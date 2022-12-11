package httpEndpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/contracts"
	"net/http"
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
	refreshToken, state, err := endpoints.Actions.RefreshTokenCreate(user)
	if err != nil {
		// Return json response
		c.JSON(500, contracts.SigninEndpointReturnDto{
			Err: "Something went wrong",
		})
		return
	}

	// Set the refresh token in the cookies
	// set sameSite to none to allow cross origin
	c.SetSameSite(http.SameSiteNoneMode)
	// set secure to true to allow only https
	c.SetCookie(
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookieName,
		refreshToken,
		60*60*24*30,
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookiePath,
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookieDomain,
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookieShouldBeSecure.Bool(),
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookieShouldBeHttpOnly.Bool())

	// Return json response
	c.JSON(200, contracts.SigninEndpointReturnDto{
		Message:     fmt.Sprintf("Welcome %s", user.Email),
		AccessToken: accessToken,
	})
}
