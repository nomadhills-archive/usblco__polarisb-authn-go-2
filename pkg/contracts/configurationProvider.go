package contracts

import (
	"encoding/json"
	"github.com/usblco/polarisb-authn-go-2/configurationSettings"
)

type ConfigurationProvider struct {
	Defaults *configurationSettings.ConfigurationSettings
	Override *configurationSettings.ConfigurationSettings
}

func NewConfigurationProvider(overrideSettings *configurationSettings.ConfigurationSettings) *ConfigurationProvider {
	cp := &ConfigurationProvider{
		Defaults: configurationSettings.NewDefaultConfigurationSettings(),
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
