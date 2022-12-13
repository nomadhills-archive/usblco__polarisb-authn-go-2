package polarisb_authn_go_2

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go-2/configurationSettings"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
	"github.com/usblco/polarisb-authn-go-2/pkg/repos"
)

type PolarisbNativeAuthnConfiguration struct {
	DatabasePathSQLLite      string                                       // The database connection string for the SQLLite database
	PolarisbUserRepo         repos.IPolarisbUserRepo                      // The base user repository for this application
	ApiPrefix                string                                       // The prefix for the API
	GinRouterGroup           *gin.RouterGroup                             // The gin router group to use for the API
	appConfigurationProvider *contracts.ConfigurationProvider             // app configuration settings
	AppConfiguration         *configurationSettings.ConfigurationSettings // app configuration settings
}
