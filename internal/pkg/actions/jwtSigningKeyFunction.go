package actions

import "github.com/usblco/polarisb-authn-go/pkg/models"

func (action *Actions) jwtSigningKeyFunction(user *models.PolarisbUser) (keyId string, signingKey string, error error) {
	return "key", "secret", nil
}
