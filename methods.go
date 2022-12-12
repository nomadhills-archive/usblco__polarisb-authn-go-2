package polarisb_authn_go_2

import (
	"github.com/usblco/polarisb-authn-go/internal"
	"github.com/usblco/polarisb-authn-go/internal/pkg/middleware"
	"github.com/usblco/polarisb-authn-go/pkg/contracts"
	"github.com/usblco/polarisb-authn-go/pkg/models"
	"time"
)

func AddPolarisbaseNativeAuthn(config *PolarisbNativeAuthnConfiguration) *PolarisbNativeAuthn {

	app := &PolarisbNativeAuthn{
		config:   config,
		internal: internal.AddPolarisbNativeAuthnInternal(),
	}

	app.initializeAppConfiguration()
	app.setupDatabaseConnection()
	app.setupPolarisbUserRepo()
	app.initializeActions()
	app.initializeMiddleware()
	app.initializeEndpoints() // This should always be called last

	return app
}

func (p *PolarisbNativeAuthn) initializeAppConfiguration() {
	p.config.appConfigurationProvider = contracts.NewConfigurationProvider(p.config.AppConfiguration)
	p.internal.InitializeAppConfiguration(p.config.appConfigurationProvider)
}

func (p *PolarisbNativeAuthn) setupDatabaseConnection() {
	if p.config.DatabasePathSQLLite != "" {
		p.internal.InitializeGormDBConnectionIfNotDoneUsingSQLite(p.config.DatabasePathSQLLite)
		return
	}

	p.internal.InitializeGormDBConnectionIfNotDoneUsingSQLite("database.db")
	return
}

func (p *PolarisbNativeAuthn) setupPolarisbUserRepo() {

	// Use the configuration to set the user repo
	if p.config.PolarisbUserRepo != nil {
		p.internal.Repos.User = p.config.PolarisbUserRepo
		return
	}

	// Use the default user repo
	// todo: handel error
	p.internal.InitializeDefaultPolarisbUserRepo()
}

func (p *PolarisbNativeAuthn) initializeActions() {
	p.internal.InitializeActions()
}

func (p *PolarisbNativeAuthn) initializeMiddleware() {
	if p.internal == nil {
		panic("internal not initialized, must call AddPolarisbaseNativeAuthn first")
	}
	if p.internal.Actions == nil {
		panic("actions not initialized")
	}
	p.internal.Middleware = middleware.InitializeMiddleware(p.internal.Actions, p.internal.AppConfiguration)
}

func (p *PolarisbNativeAuthn) initializeEndpoints() {

	if p.config.ApiPrefix != "" {
		p.internal.InitializeEndpoints(p.config.GinRouterGroup, p.config.ApiPrefix)
		return
	}

	p.internal.InitializeEndpoints(p.config.GinRouterGroup, "/authn")
}

func (p *PolarisbNativeAuthn) SetExpirationTimeFunctionAuthorizationTokens(
	function func(user *models.PolarisbUser) (expTime time.Time, error error)) *PolarisbNativeAuthn {
	p.internal.Actions.JwtFunctions.ExpirationTimeFunctionAuthorizationTokens = function
	return p
}
