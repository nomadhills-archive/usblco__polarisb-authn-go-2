package main

import (
	"github.com/gin-gonic/gin"
	polarisb_authn_go_2 "github.com/usblco/polarisb-authn-go"
)

func main() {
	// Build router
	router := gin.New()

	// create router group
	r := router.Group("")

	// Add polarisb-authn-go
	polarisb_authn_go_2.AddPolarisbaseNativeAuthn(&polarisb_authn_go_2.PolarisbNativeAuthnConfiguration{
		GinRouterGroup: r,
	})

	// Serve API (begin service loop)
	// todo: make port configurable
	// todo: handel error
	router.Run(":9096")
}
