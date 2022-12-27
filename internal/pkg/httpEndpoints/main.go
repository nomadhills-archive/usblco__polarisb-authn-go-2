package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go-2/internal/pkg/actions"
	"github.com/usblco/polarisb-authn-go-2/internal/pkg/middleware"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
)

type Endpoints struct {
	GinRouterGroup   *gin.RouterGroup
	Actions          *actions.Actions
	Middleware       *middleware.Middleware
	AppConfiguration *contracts.ConfigurationProvider
}

func InitializeEndpoints(routerGroup *gin.RouterGroup, actions *actions.Actions, middleware *middleware.Middleware, config *contracts.ConfigurationProvider, apiPrefix string) *Endpoints {
	endpoints := &Endpoints{
		Actions:          actions,
		GinRouterGroup:   routerGroup,
		Middleware:       middleware,
		AppConfiguration: config,
	}

	r := endpoints.GinRouterGroup.Group(formatApiPrefix(apiPrefix))
	authenticationRequired := r.Group("").Use(middleware.Authenticate())

	r.GET("/system/health", endpoints.HealthEndpoint)
	r.POST("/system/create-admin", endpoints.CreateAdmin)
	authenticationRequired.POST("/register", endpoints.RegisterEndpoint)
	r.POST("/signin", endpoints.SigninEndpoint)
	r.GET("/signout", endpoints.SignoutEndpoint)
	r.GET("/token/refresh", endpoints.RefreshEndpoint)
	authenticationRequired.GET("/token/info", endpoints.TokenInfoEndpoint)
	authenticationRequired.GET("/user/info", endpoints.UserInfoEndpoint)
	authenticationRequired.GET("/user/get-all", endpoints.UsersGetAllEndpoint)

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
