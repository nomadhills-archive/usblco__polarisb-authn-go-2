package polarisb_authn_go_2

import (
	"github.com/usblco/polarisb-authn-go-2/internal"
	"github.com/usblco/polarisb-authn-go-2/pkg/publicActions"
)

type PolarisbNativeAuthn struct {
	internal *internal.PolarisbNativeAuthnInternal
	config   *PolarisbNativeAuthnConfiguration
	Actions  *publicActions.PublicActions
}
