package infra

import (
	"fmt"

	"github.com/LuizEduardo-service/exe_interface/internal/entity"
)

type UserStorePostgres struct{}

func NewUserStorePostgres() *UserStorePostgres {
	return &UserStorePostgres{}
}

func (u *UserStorePostgres) Create(user *entity.User) error {

	fmt.Printf("Usuario: %q criado no banco de dados Postgres", user.Name)
	return nil
}
