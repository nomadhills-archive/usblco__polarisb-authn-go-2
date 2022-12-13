package actions

import (
	"github.com/google/uuid"
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/models"
)

// UserGetById gets a user from the database by id. Returns the user, a result state, and an error. The result state is one of the following: UserFound, UserNotFound.
func (actions *Actions) UserGetById(userId uuid.UUID) (*models.PolarisbUser, pkg.ResultState, error) {
	// look up user
	user, state, err := actions.Repos.User.LookUpUserById(nil, userId)
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
