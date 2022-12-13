package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go-2/internal/pkg/actions"
	"github.com/usblco/polarisb-authn-go-2/internal/pkg/httpEndpoints"
	"github.com/usblco/polarisb-authn-go-2/internal/pkg/middleware"
	"github.com/usblco/polarisb-authn-go-2/internal/pkg/repos"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
	"gorm.io/gorm"
)

type PolarisbNativeAuthnInternal struct {
	DB               *gorm.DB
	Repos            *repos.Repos
	Endpoints        *httpEndpoints.Endpoints
	Actions          *actions.Actions
	GinRouterGroup   *gin.RouterGroup
	Middleware       *middleware.Middleware
	AppConfiguration *contracts.ConfigurationProvider
}

func AddPolarisbNativeAuthnInternal() *PolarisbNativeAuthnInternal {
	return &PolarisbNativeAuthnInternal{
		Repos: &repos.Repos{},
	}
}
