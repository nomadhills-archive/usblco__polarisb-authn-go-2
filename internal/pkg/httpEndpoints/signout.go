package httpEndpoints

import "github.com/gin-gonic/gin"

func (endpoints *Endpoints) SignoutEndpoint(c *gin.Context) {
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
