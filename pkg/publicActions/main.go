package publicActions

import "github.com/usblco/polarisb-authn-go-2/internal"

type PublicActions struct {
	internal *internal.PolarisbNativeAuthnInternal
}

func InitializePublicActions(internal *internal.PolarisbNativeAuthnInternal) *PublicActions {
	actions := &PublicActions{
		internal: internal,
	}

	return actions
}
