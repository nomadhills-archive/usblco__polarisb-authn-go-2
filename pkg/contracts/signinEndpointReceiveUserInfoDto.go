package contracts

// SigninEndpointReceiveUserInfoDto is the struct that is used to receive the user info from the signin endpoint
type SigninEndpointReceiveUserInfoDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
