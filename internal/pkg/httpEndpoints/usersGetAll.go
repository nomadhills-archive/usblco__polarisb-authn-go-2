package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
)

func (endpoints *Endpoints) UsersGetAllEndpoint(c *gin.Context) {
	// Get all users from the database
	users, state, err := endpoints.Actions.UserGetAll()
	if state != pkg.UserFound {
		if err != nil {
			c.JSON(500, &contracts.UsersGetAllReturnDto{
				Err: err.Error(),
			})
			return
		}
		c.JSON(500, &contracts.UsersGetAllReturnDto{
			Err: "Unknown error",
		})
		return
	}

	// convert to dto
	usersDto := make([]*contracts.UserDto, len(users))
	for i, user := range users {
		usersDto[i] = &contracts.UserDto{
			Email:    user.Email,
			Id:       user.Id.String(),
			Fullname: user.Fullname,
		}
	}

	// Return the users
	c.JSON(200, &contracts.UsersGetAllReturnDto{
		Users: usersDto,
	})
}
