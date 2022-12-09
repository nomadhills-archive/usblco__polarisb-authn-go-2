package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go/pkg"
)

func (e *Endpoints) HealthEndpoint(c *gin.Context) {

	result, status, _ := e.Actions.SystemHealthInfo()
	if status != pkg.Ok {
		c.JSON(500, result)
		return
	}

	// Return json with the system info
	c.JSON(200, result)
}
