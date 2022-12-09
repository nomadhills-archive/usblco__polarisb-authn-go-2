package pkg

type ResultState int64

const (
	Ok ResultState = iota
	Unknown
	Error
	UserCreated
	UserFound
	UserNotFound
	UserAlreadyExists
	CouldNotGeneratePasswordHash
	SystemHealthReposIsNil
)

func (r ResultState) String() string {
	switch r {
	case Ok:
		return "Ok"
	case Error:
		return "Error"
	case UserCreated:
		return "UserCreated"
	default:
		return "Unknown"
	}
}
