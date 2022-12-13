package actions

import (
	"github.com/golang-jwt/jwt"
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
	"time"
)

func (actions *Actions) AccessTokenGetInfo(accessTokenString string) (*contracts.TokenInfoEndpointReturnDtoAccessTokenInfo, pkg.ResultState, error) {
	// validate access token
	accessToken, state, err := actions.AccessTokenValidate(accessTokenString)
	if state != pkg.JwtTokenVerified {
		return nil, state, err
	}

	accessTokenInfo := &contracts.TokenInfoEndpointReturnDtoAccessTokenInfo{}

	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	accessTokenExpirationTime := accessTokenClaims["exp"].(float64)
	accessTokenInfo.AccessTokenExpiresIn = int64(accessTokenExpirationTime) - time.Now().Unix()
	accessTokenInfo.AccessTokenString = accessTokenString
	accessTokenInfo.AccessTokenValid = accessToken.Valid

	return accessTokenInfo, pkg.Ok, nil
}
