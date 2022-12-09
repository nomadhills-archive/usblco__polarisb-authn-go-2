package actions

import "github.com/usblco/polarisb-authn-go/internal/pkg/repos"

type Actions struct {
	Repos *repos.Repos
}

func InitializeActions(repos *repos.Repos) *Actions {
	return &Actions{
		Repos: repos,
	}
}
