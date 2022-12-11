package polarisb_authn_go_2

import (
	"github.com/usblco/polarisb-authn-go/internal"
)

type PolarisbNativeAuthn struct {
	internal *internal.PolarisbNativeAuthnInternal
	config   *PolarisbNativeAuthnConfiguration
}
