package infra

import "github.com/LuizEduardo-service/exe_interface/internal/entity"

type UserStoreInterface interface {
	Create(user *entity.User) error
}
