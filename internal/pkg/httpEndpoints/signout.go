package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (endpoints *Endpoints) SignoutEndpoint(c *gin.Context) {
	// Set the refresh token in the cookies
	// set sameSite to none to allow cross origin
	c.SetSameSite(http.SameSiteNoneMode)
	// Remove the refresh token from the cookies
	c.SetCookie(
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookieName,
		"", -1,
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookiePath,
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookieDomain,
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookieShouldBeSecure.Bool(),
		endpoints.AppConfiguration.Defaults.RefreshTokenSettings.CookieShouldBeHttpOnly.Bool())
	// Return a 200 response
	c.String(200, "ok")
}
