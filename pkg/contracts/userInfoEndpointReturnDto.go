package contracts

type UserInfoEndpointReturnDto struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Err           string `json:"err,omitempty"`
}
