package contracts

type RefreshEndpointReturnDto struct {
	Err         string `json:"err,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}
