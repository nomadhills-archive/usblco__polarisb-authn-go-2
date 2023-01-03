package publicActions

import (
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/models"
)

func (actions *PublicActions) CheckIfAuthorizedByRole(role string, user *models.PolarisbUser) (bool, pkg.ResultState, error) {
	return actions.internal.Actions.CheckIfAuthorizedByRole(role, user)
}
