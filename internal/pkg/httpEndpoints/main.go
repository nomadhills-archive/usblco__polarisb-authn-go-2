package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go/internal/pkg/actions"
)

type Endpoints struct {
	GinRouterGroup *gin.RouterGroup
	Actions        *actions.Actions
}

func InitializeEndpoints(routerGroup *gin.RouterGroup, actions *actions.Actions, apiPrefix string) *Endpoints {
	endpoints := &Endpoints{
		Actions:        actions,
		GinRouterGroup: routerGroup,
	}

	r := endpoints.GinRouterGroup.Group(formatApiPrefix(apiPrefix))
	r.GET("/system/health", endpoints.HealthEndpoint)
	r.POST("/register", endpoints.RegisterEndpoint)
	r.POST("/signin", endpoints.SigninEndpoint)
	endpoints.GinRouterGroup = r

	return endpoints
}

func formatApiPrefix(apiPrefix string) string {
	if apiPrefix == "" {
		return ""
	}
	if apiPrefix[0] != '/' {
		apiPrefix = "/" + apiPrefix
	}
	if apiPrefix[len(apiPrefix)-1] != '/' {
		apiPrefix = apiPrefix + "/"
	}
	return apiPrefix
}
