package contracts

type SystemHealthEndpointReturnDto struct {
	Err     string `json:"err,omitempty"`
	Message string `json:"message,omitempty"`
}
