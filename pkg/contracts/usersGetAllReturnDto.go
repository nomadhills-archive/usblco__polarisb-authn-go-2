package contracts

type UsersGetAllReturnDto struct {
	Err   string     `json:"err,omitempty"`
	Users []*UserDto `json:"users"`
}
