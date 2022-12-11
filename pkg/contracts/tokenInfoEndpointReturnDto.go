package contracts

type TokenInfoEndpointReturnDto struct {
	Err                   string                                     `json:"err,omitempty"`
	CurrentUserId         string                                     `json:"currentUserId,omitempty"`
	AccessTokenInfo       *TokenInfoEndpointReturnDtoAccessTokenInfo `json:"accessTokenInfo,omitempty"`
	RefreshTokenValid     bool                                       `json:"refreshTokenValid,omitempty"`
	RefreshTokenExpiresAt int64                                      `json:"refreshTokenExpiresAt,omitempty"`
}

type TokenInfoEndpointReturnDtoAccessTokenInfo struct {
	AccessTokenValid     bool   `json:"accessTokenValid,omitempty"`
	AccessTokenExpiresIn int64  `json:"accessTokenExpiresAtIn,omitempty"`
	AccessTokenString    string `json:"accessToken,omitempty"`
}

type TokenInfoEndpointReturnDtoRefreshTokenInfo struct {
	RefreshTokenValid     bool  `json:"accessTokenValid,omitempty"`
	RefreshTokenExpiresIn int64 `json:"accessTokenExpiresAtIn,omitempty"`
}
