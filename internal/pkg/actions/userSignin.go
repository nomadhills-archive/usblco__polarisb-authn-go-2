package actions

import (
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/models"
)

// UserSigninByEmail signs in a user by email and password and returns a user object and result state. Result state should be UserSigninSuccess for success or InvalidPassword or UserNotFound.
func (actions *Actions) UserSigninByEmail(email string, password string) (*models.PolarisbUser, pkg.ResultState, error) {
	// look up user
	user, state, err := actions.UserGetByEmail(email)
	if state == pkg.UserNotFound {
		return nil, pkg.UserNotFound, nil
	}
	if state != pkg.UserFound {
		return nil, state, err
	}

	// Check if password is correct
	if !user.CheckPassword(password) {
		return nil, pkg.InvalidPasswd, nil
	}

	return user, pkg.UserSigninSuccess, nil
}
