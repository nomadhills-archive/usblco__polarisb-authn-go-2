package actions

import (
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/models"
)

func (actions *Actions) UserGetByEmail(email string) (*models.PolarisbUser, pkg.ResultState, error) {
	// look up user
	user, state, err := actions.Repos.User.LookUpUserByEmail(nil, email)
	if err != nil {
		return nil, state, err
	}

	if state == pkg.UserNotFound {
		return nil, pkg.UserNotFound, nil
	}

	if state == pkg.UserFound {
		return user, pkg.UserFound, nil
	}

	return nil, pkg.Unknown, err
}
