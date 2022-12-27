package actions

import (
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/models"
	"strings"
)

func (actions *Actions) CheckIfAuthorizedByRole(role string, user *models.PolarisbUser) (bool, pkg.ResultState, error) {

	// split role string by ,
	roles := strings.Split(user.Role, ",")
	// loop through roles
	for _, currentRole := range roles {
		// check if user has role
		if role == currentRole {
			// return true
			return true, pkg.UserHasRole, nil
		}
	}
	// return false
	return false, pkg.UserDoesNotHaveRole, nil
}
