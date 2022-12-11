package polarisb_authn_go_2

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go/ConfigurationSettings"
	"github.com/usblco/polarisb-authn-go/pkg/contracts"
	"github.com/usblco/polarisb-authn-go/pkg/repos"
)

type PolarisbNativeAuthnConfiguration struct {
	DatabasePathSQLLite      string                                       // The database connection string for the SQLLite database
	PolarisbUserRepo         repos.IPolarisbUserRepo                      // The base user repository for this application
	ApiPrefix                string                                       // The prefix for the API
	GinRouterGroup           *gin.RouterGroup                             // The gin router group to use for the API
	appConfigurationProvider *contracts.ConfigurationProvider             // app configuration settings
	AppConfiguration         *ConfigurationSettings.ConfigurationSettings // app configuration settings
}
