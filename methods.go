package polarisb_authn_go_2

import "github.com/usblco/polarisb-authn-go/internal"

func AddPolarisbaseNativeAuthn(config *PolarisbNativeAuthnConfiguration) *PolarisbNativeAuthn {

	app := &PolarisbNativeAuthn{
		config:   config,
		internal: internal.AddPolarisbNativeAuthnInternal(),
	}

	app.setupDatabaseConnection()
	app.setupPolarisbUserRepo()
	app.initializeActions()
	app.initializeEndpoints() // This should always be called last

	return app
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

func (p *PolarisbNativeAuthn) initializeEndpoints() {

	if p.config.ApiPrefix != "" {
		p.internal.InitializeEndpoints(p.config.GinRouterGroup, p.config.ApiPrefix)
		return
	}

	p.internal.InitializeEndpoints(p.config.GinRouterGroup, "/api")
}
