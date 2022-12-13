package actions

import "github.com/usblco/polarisb-authn-go-2/pkg"

// UserDoesUserAlreadyExistsByEmail looks up a user by email, return state should be UserFound for success
func (action *Actions) UserDoesUserAlreadyExistsByEmail(address string) (pkg.ResultState, error) {
	// Look up user by email
	_, state, err := action.Repos.User.LookUpUserByEmail(nil, address)
	if err != nil {
		return pkg.Error, err
	}

	if state == pkg.UserFound {
		return pkg.UserFound, nil
	}

	if state == pkg.UserNotFound {
		return pkg.UserNotFound, nil
	}

	// Yes, user exists
	return pkg.Unknown, nil
}
