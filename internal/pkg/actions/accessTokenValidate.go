package actions

import (
	"github.com/golang-jwt/jwt"
	"github.com/usblco/polarisb-authn-go/pkg"
	"time"
)

// AccessTokenValidate validates an access token. If token is vaalid, it returns a result state of JwtTokenVerified. If token is invalid, it returns a result state of JwtTokenInvalid.
func (actions *Actions) AccessTokenValidate(accessTokenString string) (refreshToken *jwt.Token, state pkg.ResultState, err error) {

	// parse token
	token, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
		// validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidJwtSigningMethod
		}

		// validate key id
		if token.Header["kid"] == "" {
			return nil, errInvalidJwtKeyId
		}

		// get signing key from database by key id
		signingKey, err := actions.jwtLookupSigningKeyFunction(token.Header["kid"].(string))
		if err != nil {
			return nil, errInvalidJwtToken
		}

		// validate token type
		if token.Claims.(jwt.MapClaims)["type"] != "authn" {
			return nil, errInvalidJwtTokenType
		}

		// validate token is not expired
		if time.Now().Unix() > int64(token.Claims.(jwt.MapClaims)["exp"].(float64)) {
			return nil, errJwtTokenExpired
		}

		return []byte(signingKey), nil
	})

	var resultState = translateErrorToResultState(err)

	if resultState != pkg.JwtTokenVerified {
		return nil, resultState, err
	}

	if resultState == pkg.JwtTokenVerified {
		return token, resultState, nil
	}

	return nil, pkg.Unknown, nil
}

