package actions

import (
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/models"
)

func (actions *Actions) UserIsEmailVerified(user *models.PolarisbUser) pkg.ResultState {
	if user.VerifiedEmail == "true" {
		return pkg.UserEmailVerified
	} else {
		return pkg.UserEmailNotVerified
	}
}
