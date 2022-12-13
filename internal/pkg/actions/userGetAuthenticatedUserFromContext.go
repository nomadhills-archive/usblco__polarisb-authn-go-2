package actions

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/models"
)

func (actions *Actions) UserGetAuthenticatedUserFromContext(c *gin.Context) (user *models.PolarisbUser, state pkg.ResultState, error error) {
	// If we get here, the user is authenticated
	// Get the user id from the context
	userId, ok := c.Get("current_user")
	if !ok {
		return nil, pkg.UserNotFound, errors.New("Current user not found in context")
	}

	// convert the user id to a google uuid
	userUuid, err := uuid.Parse(userId.(string))
	if err != nil {
		return nil, pkg.UserInvalidId, err
	}

	// Get the user from the database
	user, state, err = actions.Repos.User.LookUpUserById(nil, userUuid)
	if state != pkg.UserFound {
		return nil, state, err
	}

	if state == pkg.UserFound {
		return user, pkg.UserAuthenticated, nil
	}

	return user, pkg.UserNotAuthenticated, err
}
