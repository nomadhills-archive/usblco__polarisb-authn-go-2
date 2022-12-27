package repos

import (
	"context"
	"github.com/google/uuid"
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/models"
)

type IPolarisbUserRepo interface {
	// CreateUserObject creates a new user object, return state should be UserCreated for success
	CreateUserObject(ctx context.Context, email string, passwordHash string) (*models.PolarisbUser, pkg.ResultState, error)

	// LookUpUserByEmail looks up a user by email, return state should be UserFound for success and UserNotFound for fail
	LookUpUserByEmail(ctx context.Context, email string) (*models.PolarisbUser, pkg.ResultState, error)

	// LookUpUserById looks up a user by id, return state should be UserFound for success
	LookUpUserById(ctx context.Context, id uuid.UUID) (*models.PolarisbUser, pkg.ResultState, error)

	// GetAll returns all users, return state should be UserFound for success
	GetAll() ([]*models.PolarisbUser, pkg.ResultState, error)

	// GetFirst returns the first user, return state should be UserFound for success
	GetFirst() (*models.PolarisbUser, pkg.ResultState, error)

	// UpdateUserObject updates a user object, return state should be UserUpdated for success
	UpdateUserObject(ctx context.Context, user *models.PolarisbUser) (pkg.ResultState, error)
}
