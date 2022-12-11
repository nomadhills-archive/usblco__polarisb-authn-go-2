package actions

import (
	"errors"
	"github.com/usblco/polarisb-authn-go/pkg"
)

var (
	errInvalidJwtSigningMethod = errors.New("invalid jwt signing method")
	errInvalidJwtKeyId         = errors.New("invalid jwt key id")
	errInvalidJwtToken         = errors.New("invalid jwt token")
	errInvalidJwtTokenType     = errors.New("invalid jwt token type")
	errJwtTokenExpired         = errors.New("jwt token expired")
	jwtTokenVerified           = errors.New("jwt token verified")
)

// translate error to result state
func translateErrorToResultState(err error) pkg.ResultState {

	if err == nil {
		return pkg.JwtTokenVerified
	}

	switch err.Error() {
	case errInvalidJwtSigningMethod.Error():
		//fmt.Printf("translateErrorToResultState(); translated to pkg.InvalidJwtSigningMethod \n")
		return pkg.InvalidJwtSigningMethod
	case errInvalidJwtKeyId.Error():
		//fmt.Printf("translateErrorToResultState(); translated to pkg.InvalidJwtKeyId \n")
		return pkg.InvalidJwtKeyId
	case errInvalidJwtToken.Error():
		//fmt.Printf("translateErrorToResultState(); translated to pkg.InvalidJwtToken \n")
		return pkg.InvalidJwtToken
	case errInvalidJwtTokenType.Error():
		//fmt.Printf("translateErrorToResultState(); translated to pkg.InvalidJwtTokenType \n")
		return pkg.InvalidJwtTokenType
	case errJwtTokenExpired.Error():
		//fmt.Printf("translateErrorToResultState(); translated to pkg.JwtTokenExpired \n")
		return pkg.JwtTokenExpired
	case jwtTokenVerified.Error():
		//fmt.Printf("translateErrorToResultState(); translated to pkg.JwtTokenVerified \n")
		return pkg.JwtTokenVerified

	}
	return pkg.Unknown
}
