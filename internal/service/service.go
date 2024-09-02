package service

import (
	"github.com/LuizEduardo-service/exe_interface/internal/entity"
	"github.com/LuizEduardo-service/exe_interface/internal/infra"
)

type UserService struct {
	userStore infra.UserStoreInterface
}

func NewUserService(userStore infra.UserStoreInterface) (*UserService, error) {
	return &UserService{
		userStore: userStore,
	}, nil
}

func (u *UserService) CreateUserHandler(name string) {
	user := entity.NewUser(name)
	u.userStore.Create(user)
}
