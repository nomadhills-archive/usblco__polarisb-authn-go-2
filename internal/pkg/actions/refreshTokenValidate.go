package actions

import (
	"github.com/golang-jwt/jwt"
	"github.com/usblco/polarisb-authn-go/pkg"
)

func (actions *Actions) RefreshTokenValidate(refreshTokenString string) (refreshToken *jwt.Token, state pkg.ResultState, error error) {
	// Parse the token
	
}
