package middleware

import (
	"github.com/usblco/polarisb-authn-go-2/internal/pkg/actions"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
)

type Middleware struct {
	actions          *actions.Actions
	appConfiguration *contracts.ConfigurationProvider
}

func InitializeMiddleware(actions *actions.Actions, config *contracts.ConfigurationProvider) *Middleware {
	return &Middleware{
		actions:          actions,
		appConfiguration: config,
	}
}
