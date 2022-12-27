package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
)

func (endpoint *Endpoints) CreateAdmin(c *gin.Context) {
	// check if this is the firstUser
	isFirstUser := false
	_, state, _ := endpoint.Actions.Repos.User.GetFirst()
	if state == pkg.UserNotFound {
		isFirstUser = true
	}

	// if this is the first user, create the admin user
	if isFirstUser {
		// create the admin user
		_, state, _ := endpoint.Actions.UserCreate("admin@admin.admin", "admin")
		if state != pkg.UserCreated {
			if state == pkg.UserAlreadyExists {
				// Return json response
				c.JSON(409, contracts.RegisterEndpointReturnDto{
					Err: "user already exists",
				})
				return
			}
			// Return json response
			c.JSON(500, contracts.RegisterEndpointReturnDto{
				Err: "user could not be created",
			})
			return
		}
	} else {
		// Return json response
		c.JSON(401, contracts.RegisterEndpointReturnDto{
			Err: "not authorized, only the first user can be the admin user using this method",
		})
	}
}
