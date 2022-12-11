package actions

import (
	"github.com/usblco/polarisb-authn-go/pkg/models"
	"time"
)

type JwtFunctions struct {
	ExpirationTimeFunctionAuthorizationTokens func(user *models.PolarisbUser) (expTime time.Time, error error)
	ExpirationTimeFunctionRefreshTokens       func(user *models.PolarisbUser) (expTime time.Time, error error)
	SigningKeyFunction                        func(user *models.PolarisbUser) (keyId string, signingKey string, error error)
	LookupSigningKeyFunction                  func(kid string) (signingKey string, error error)
}

func newJwtFunctions(actions *Actions) *JwtFunctions {
	return &JwtFunctions{
		ExpirationTimeFunctionAuthorizationTokens: actions.jwtExpirationTimeFunctionAccessTokens,
		ExpirationTimeFunctionRefreshTokens:       actions.jwtExpirationTimeFunctionRefreshTokens,
		SigningKeyFunction:                        actions.jwtSigningKeyFunction,
		LookupSigningKeyFunction:                  actions.jwtLookupSigningKeyFunction,
	}
}
