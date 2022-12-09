package actions

import (
	"github.com/golang-jwt/jwt"
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/models"
)

func (actions *Actions) AccessTokenCreate(user *models.PolarisbUser) (accessToken string, state pkg.ResultState, error error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// get singning key
	keyId, signingKey, err := actions.JwtFunctions.SigningKeyFunction(user)
	if err != nil {
		return "", pkg.CouldNotGeJwtSigningKey, err
	}

	// get expiration time
	expTime, err := actions.JwtFunctions.ExpirationTimeFunctionAuthorizationTokens(user)
	if err != nil {
		return "", pkg.CouldNotGetJwtExpirationTime, err
	}

	// set key id
	token.Header["kid"] = keyId

	// Set some claims
	token.Claims = jwt.MapClaims{
		"type":     "authn",
		"id":       user.Id,
		"email":    user.Email,
		"username": user.Username,
		"fullname": user.Fullname,
		"role":     user.Role,
		"exp":      expTime.Unix(),
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", pkg.CouldNotSignJwt, err
	}

	return tokenString, pkg.AccessTokenCreated, nil
}
