package actions

import (
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/models"
)

func (actions *Actions) UserGetAll() ([]*models.PolarisbUser, pkg.ResultState, error) {
	result, state, err := actions.Repos.User.GetAll()
	return result, state, err
}
