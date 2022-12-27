package contracts

type UserDto struct {
	Email     string `json:"email"`
	Id        string `json:"id"`
	Fullname  string `json:"fullname"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
