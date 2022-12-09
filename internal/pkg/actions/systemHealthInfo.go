package actions

import (
	"errors"
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/contracts"
)

func (action *Actions) SystemHealthInfo() (contracts.SystemHealthEndpointReturnDto, pkg.ResultState, error) {
	if action.Repos == nil {
		return contracts.SystemHealthEndpointReturnDto{
			Message: "Repos is nil",
		}, pkg.SystemHealthReposIsNil, errors.New("Repos is nil")
	}

	return contracts.SystemHealthEndpointReturnDto{
		Message: "Polarisb Native Authn is up and running!",
	}, pkg.Ok, nil
}
