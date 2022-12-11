package actions

import (
	"github.com/usblco/polarisb-authn-go/pkg/models"
	"time"
)

func (action *Actions) jwtExpirationTimeFunctionAccessTokens(user *models.PolarisbUser) (expTime time.Time, error error) {
	return time.Now().Add(time.Minute * 2), nil
}
