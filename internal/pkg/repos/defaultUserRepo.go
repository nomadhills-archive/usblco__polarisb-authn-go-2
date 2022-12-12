package repos

import (
	"context"
	"github.com/google/uuid"
	"github.com/usblco/polarisb-authn-go/pkg"
	"github.com/usblco/polarisb-authn-go/pkg/models"
	"github.com/usblco/polarisb-authn-go/pkg/repos"
	"gorm.io/gorm"
	"strings"
)

type UsersStore struct {
	dbconn *gorm.DB
}

// InitializeBasicUsersStore new project store
func InitializeUsersRepo(dbconn *gorm.DB) repos.IPolarisbUserRepo {
	// todo: handle error
	dbconn.AutoMigrate(&models.PolarisbUser{})
	return &UsersStore{dbconn: dbconn}
}

func (k *UsersStore) CreateUserObject(ctx context.Context, email string, passwordHash string) (*models.PolarisbUser, pkg.ResultState, error) {
	// User object
	user := models.PolarisbUser{
		Id:           uuid.New(),
		Email:        strings.ToLower(email),
		Passwordhash: passwordHash,
	}

	// Create
	k.dbconn.Create(&user)
	return &user, pkg.UserCreated, nil
}

func (k *UsersStore) LookUpUserByEmail(ctx context.Context, email string) (*models.PolarisbUser, pkg.ResultState, error) {
	// User object
	result := &models.PolarisbUser{}

	k.dbconn.Model(models.PolarisbUser{Email: strings.ToLower(email)}).First(result)

	if result.Email != strings.ToLower(email) {
		return nil, pkg.UserNotFound, nil
	}

	return result, pkg.UserFound, nil
}

func (k *UsersStore) LookUpUserById(ctx context.Context, id uuid.UUID) (*models.PolarisbUser, pkg.ResultState, error) {
	// User object
	result := &models.PolarisbUser{}

	k.dbconn.Model(models.PolarisbUser{Id: id}).First(result)

	if result.Id != id {
		return nil, pkg.UserNotFound, nil
	}

	return result, pkg.UserFound, nil
}

func (k *UsersStore) UpdateUserObject(ctx context.Context, user *models.PolarisbUser) (pkg.ResultState, error) {
	k.dbconn.Save(user)
	return pkg.UserUpdated, nil
}
