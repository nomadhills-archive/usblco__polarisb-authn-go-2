package polarisb_authn_go_2

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/usblco/polarisb-authn-go-2/pkg"
)

func (p *PolarisbNativeAuthn) ValidateAccessToken(accessTokenString string) (refreshToken *jwt.Token, state pkg.ResultState, err error) {
	return p.internal.Actions.AccessTokenValidate(accessTokenString)
}

func (p *PolarisbNativeAuthn) AuthenticateMiddleware() gin.HandlerFunc {
	return p.internal.Middleware.Authenticate()
}
