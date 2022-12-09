package actions

func (action *Actions) jwtLookupSigningKeyFunction(kid string) (signingKey string, error error) {
	return "secret", nil
}
