package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/usblco/polarisb-authn-go/internal/pkg/actions"
	"github.com/usblco/polarisb-authn-go/internal/pkg/httpEndpoints"
	"github.com/usblco/polarisb-authn-go/internal/pkg/repos"
	"gorm.io/gorm"
)

func (internal *PolarisbNativeAuthnInternal) InitializeGormDBConnectionIfNotDoneUsingSQLite(sqlitePath string) {
	if internal.DB != nil {
		return
	}

	// init db connection
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	internal.DB = db
	return
}

func (internal *PolarisbNativeAuthnInternal) InitializeDefaultPolarisbUserRepo() error {
	if internal.DB == nil {
		panic("failed to connect database")
	}

	// init user repo
	internal.Repos.User = repos.InitializeUsersRepo(internal.DB)

	return nil
}

func (internal *PolarisbNativeAuthnInternal) InitializeActions() {
	internal.Actions = actions.InitializeActions(internal.Repos)
}

func (internal *PolarisbNativeAuthnInternal) InitializeEndpoints(routerGroup *gin.RouterGroup, apiPrefix string) {
	if internal.Actions == nil {
		panic("actions not initialized")
	}
	internal.GinRouterGroup = routerGroup
	internal.Endpoints = httpEndpoints.InitializeEndpoints(routerGroup, internal.Actions, apiPrefix)
}
