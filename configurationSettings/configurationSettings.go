package configurationSettings

type ConfigurationSettings struct {
	RefreshTokenSettings *RefreshTokenSettings `json:"refreshTokenSettings,omitempty"`
}

func NewDefaultConfigurationSettings() *ConfigurationSettings {
	return &ConfigurationSettings{
		RefreshTokenSettings: NewRefreshTokenSettings(),
	}
}
