package repos

import (
	"context"
	"github.com/google/uuid"
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/models"
)

type IPolarisbUserRepo interface {
	// CreateUserObject creates a new user object, return state should be UserCreated for success
	CreateUserObject(ctx context.Context, email string, passwordHash string) (*models.PolarisbUser, pkg.ResultState, error)

	// LookUpUserByEmail looks up a user by email, return state should be UserFound for success and UserNotFound for fail
	LookUpUserByEmail(ctx context.Context, email string) (*models.PolarisbUser, pkg.ResultState, error)

	// LookUpUserById looks up a user by id, return state should be UserFound for success
	LookUpUserById(ctx context.Context, id uuid.UUID) (*models.PolarisbUser, pkg.ResultState, error)
}
