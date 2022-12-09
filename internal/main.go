package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go/internal/pkg/actions"
	"github.com/usblco/polarisb-authn-go/internal/pkg/httpEndpoints"
	"github.com/usblco/polarisb-authn-go/internal/pkg/repos"
	"gorm.io/gorm"
)

type PolarisbNativeAuthnInternal struct {
	DB             *gorm.DB
	Repos          *repos.Repos
	Endpoints      *httpEndpoints.Endpoints
	Actions        *actions.Actions
	GinRouterGroup *gin.RouterGroup
}

func AddPolarisbNativeAuthnInternal() *PolarisbNativeAuthnInternal {
	return &PolarisbNativeAuthnInternal{
		Repos: &repos.Repos{},
	}
}
