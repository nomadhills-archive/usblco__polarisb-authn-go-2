package actions

import (
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/crypto"
	"github.com/usblco/polarisb-authn-go-2/pkg/models"
)

func (actions *Actions) UserNewVerifyEmailCode(user *models.PolarisbUser) (string, pkg.ResultState, error) {
	// generate a secret code using crypto/rand
	secureCode, err := crypto.GenerateRandomStringURLSafe(32)
	if err != nil {
		return "", pkg.UserCouldNotCreateVerifyEmailCode, err
	}

	// save the secret code to the user
	user.VerifiedEmailCode = secureCode

	// save the user
	state, err := actions.Repos.User.UpdateUserObject(nil, user)
	if state != pkg.UserUpdated {
		return "", pkg.UserCouldNotCreateVerifyEmailCode, err
	}

	return secureCode, pkg.UserCreateVerifyEmailCode, nil
}
