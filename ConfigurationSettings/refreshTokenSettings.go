package ConfigurationSettings

type RefreshTokenSettings struct {
	CookieShouldBeHttpOnly BoolExtended `json:"cookieShouldBeHttpOnly,omitempty"`
	CookieShouldBeSecure   BoolExtended `json:"cookieShouldBeSecure,omitempty"`
	CookieName             string       `json:"cookieName,omitempty"`
	CookiePath             string       `json:"cookiePath,omitempty"`
	CookieDomain           string       `json:"cookieDomain,omitempty"`
}

func NewRefreshTokenSettings() *RefreshTokenSettings {
	return &RefreshTokenSettings{
		CookieShouldBeHttpOnly: True,
		CookieShouldBeSecure:   True,
		CookieName:             "rft",
		CookiePath:             "/",
		CookieDomain:           "",
	}
}
