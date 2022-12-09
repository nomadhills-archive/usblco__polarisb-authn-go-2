package actions

import (
	"github.com/usblco/polarisb-authn-go/internal/pkg/repos"
)

type Actions struct {
	Repos        *repos.Repos
	JwtFunctions *JwtFunctions
}

func InitializeActions(repos *repos.Repos) *Actions {
	actions := &Actions{
		Repos:        repos,
		JwtFunctions: &JwtFunctions{},
	}

	actions.JwtFunctions = newJwtFunctions(actions)

	return actions
}
