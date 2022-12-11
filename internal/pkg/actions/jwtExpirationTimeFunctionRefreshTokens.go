package actions

import (
	"github.com/usblco/polarisb-authn-go/pkg/models"
	"time"
)

func (action *Actions) jwtExpirationTimeFunctionRefreshTokens(user *models.PolarisbUser) (expTime time.Time, error error) {
	return time.Now().Add(time.Hour * 24 * 1), nil
}
