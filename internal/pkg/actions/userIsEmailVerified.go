package actions

import (
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/models"
)

func (actions *Actions) UserIsEmailVerified(user *models.PolarisbUser) pkg.ResultState {
	if user.VerifiedEmail == "true" {
		return pkg.UserEmailVerified
	} else {
		return pkg.UserEmailNotVerified
	}
}
