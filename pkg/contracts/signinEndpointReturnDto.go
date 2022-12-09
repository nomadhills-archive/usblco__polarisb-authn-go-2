package contracts

// SigninEndpointReturnDto is the struct that is used to return the user info from the signin endpoint
type SigninEndpointReturnDto struct {
	Username    string `json:"username,omitempty"`
	Err         string `json:"err,omitempty"`
	Message     string `json:"message,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}
