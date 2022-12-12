package main

import (
	"github.com/gin-gonic/gin"
	polarisb_authn_go_2 "github.com/usblco/polarisb-authn-go"
	"github.com/usblco/polarisb-authn-go/configurationSettings"
	"github.com/usblco/polarisb-authn-go/pkg/models"
	"time"
)

func main() {
	// Build router
	router := gin.New()

	// create router group
	r := router.Group("")

	// Add polarisb-authn-go
	polarisb_authn_go_2.AddPolarisbaseNativeAuthn(&polarisb_authn_go_2.PolarisbNativeAuthnConfiguration{
		GinRouterGroup: r,
		AppConfiguration: &configurationSettings.ConfigurationSettings{
			RefreshTokenSettings: &configurationSettings.RefreshTokenSettings{
				CookieShouldBeSecure:   configurationSettings.False,
				CookieShouldBeHttpOnly: configurationSettings.False,
			},
		},
	}).SetExpirationTimeFunctionAuthorizationTokens(func(user *models.PolarisbUser) (expTime time.Time, error error) {
		return time.Now().Add(time.Second * 45), nil
	})

	// Serve API (begin service loop)
	// todo: make port configurable
	// todo: handel error
	router.Run(":9096")
}
