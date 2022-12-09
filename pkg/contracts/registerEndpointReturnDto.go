package contracts

type RegisterEndpointReturnDto struct {
	Err     string `json:"err,omitempty"`
	Message string `json:"message,omitempty"`
}
