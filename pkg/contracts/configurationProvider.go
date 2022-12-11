package contracts

import (
	"encoding/json"
	"github.com/usblco/polarisb-authn-go/ConfigurationSettings"
)

type ConfigurationProvider struct {
	Defaults *ConfigurationSettings.ConfigurationSettings
	Override *ConfigurationSettings.ConfigurationSettings
}

func NewConfigurationProvider(overrideSettings *ConfigurationSettings.ConfigurationSettings) *ConfigurationProvider {
	cp := &ConfigurationProvider{
		Defaults: ConfigurationSettings.NewDefaultConfigurationSettings(),
		Override: overrideSettings,
	}

	cp.WriteOverrideSettings()

	return cp
}

func (cp *ConfigurationProvider) WriteOverrideSettings() {
	if cp.Override == nil {
		return
	}

	// convert Override to a json string
	overrideAsJson, err := json.Marshal(cp.Override)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(overrideAsJson, cp.Defaults)
}
