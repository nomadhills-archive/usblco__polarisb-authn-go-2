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
	UserAuthenticated
	UserNotAuthenticated
	UserInvalidId
	UserEmailVerified
	UserEmailNotVerified
	InvalidPasswd
	UserSigninSuccess
	CouldNotGeneratePasswordHash
	CouldNotGeJwtSigningKey
	CouldNotGetJwtExpirationTime
	CouldNotSignJwt
	InvalidJwtSigningMethod
	InvalidJwtKeyId
	InvalidJwtToken
	InvalidJwtTokenType
	JwtTokenExpired
	JwtTokenVerified
	AccessTokenCreated
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
