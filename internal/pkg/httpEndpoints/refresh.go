package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
)

func (endpoints *Endpoints) RefreshEndpoint(c *gin.Context) {
	// Get refresh token from cookies
	refreshTokenString, err := getRefreshTokenFromCookies(c)
	if err != nil {
		c.JSON(401, &contracts.RefreshEndpointReturnDto{
			Err: "Refresh token not found",
		})
		return
	}

	// Validate the refresh token
	refreshToken, state, err := endpoints.Actions.RefreshTokenValidate(refreshTokenString)
	if state != pkg.JwtTokenVerified {
		c.JSON(401, &contracts.RefreshEndpointReturnDto{
			Err: "Refresh token not valid",
		})
		return
	}
	// check validity of the refresh token again
	if !refreshToken.Valid {
		c.JSON(401, &contracts.RefreshEndpointReturnDto{
			Err: "Refresh token not valid",
		})
		return
	}

	// get claims as a map
	claims := refreshToken.Claims.(jwt.Claims)

	// get the user id from the claims
	userId := claims.(jwt.MapClaims)["id"].(string)

	// convert the user id to a uuid
	userIdUuid, err := uuid.Parse(userId)
	if err != nil {
		c.JSON(401, &contracts.RefreshEndpointReturnDto{
			Err: "Refresh token not valid",
		})
		return
	}

	// get the user from the database
	user, state, err := endpoints.Actions.UserGetById(userIdUuid)
	if state != pkg.UserFound {
		c.JSON(401, &contracts.RefreshEndpointReturnDto{
			Err: "User not found",
		})
		return
	}

	// generate a new access token
	accessToken, state, err := endpoints.Actions.AccessTokenCreate(user)
	if state != pkg.AccessTokenCreated {
		c.JSON(401, &contracts.RefreshEndpointReturnDto{
			Err: "Access token not created",
		})
		return
	}

	// return json response
	c.JSON(200, &contracts.RefreshEndpointReturnDto{
		AccessToken: accessToken,
	})
}
