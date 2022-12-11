package actions

import (
	"github.com/usblco/polarisb-authn-go/internal/pkg/repos"
	"github.com/usblco/polarisb-authn-go/pkg/contracts"
)

type Actions struct {
	Repos            *repos.Repos
	JwtFunctions     *JwtFunctions
	AppConfiguration *contracts.ConfigurationProvider
}

func InitializeActions(repos *repos.Repos, config *contracts.ConfigurationProvider) *Actions {
	actions := &Actions{
		Repos:            repos,
		JwtFunctions:     &JwtFunctions{},
		AppConfiguration: config,
	}

	actions.JwtFunctions = newJwtFunctions(actions)

	return actions
}
