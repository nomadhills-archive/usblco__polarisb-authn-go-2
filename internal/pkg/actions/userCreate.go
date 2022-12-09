package actions

import (
	"github.com/usblco/polarisb-authn-go/internal/pkg/internalModels"
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/models"
)

// UserCreate creates a new user. Returns the user object, the result state, and an error. The result state should be UserCreated for success.
// error states are: UserAlreadyExists, CouldNotGeneratePasswordHash
func (actions *Actions) UserCreate(emailAddress string, password string) (*models.PolarisbUser, pkg.ResultState, error) {
	// validate user does not exist
	doesUserExistStatus, _ := actions.UserDoesUserAlreadyExistsByEmail(emailAddress)
	if doesUserExistStatus == pkg.UserFound {
		return &models.PolarisbUser{}, pkg.UserAlreadyExists, nil
	}

	// Create new basePasswordObject
	passwordObj := &internalModels.PasswordObject{}

	// Begin change password process and generate hash
	err := passwordObj.GenerateAndSetPasswordHash(password)
	if err != nil {
		return nil, pkg.CouldNotGeneratePasswordHash, err
	}

	// get password object as base64 string
	passwordObjBase64, err := passwordObj.EncodeToBase64()

	// Create user object
	user, state, err := actions.Repos.User.CreateUserObject(nil, emailAddress, passwordObjBase64)
	if state != pkg.UserCreated {
		return nil, state, err
	}

	return user, pkg.UserCreated, nil
}
