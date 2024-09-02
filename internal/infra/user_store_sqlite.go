package infra

import (
	"fmt"

	"github.com/LuizEduardo-service/exe_interface/internal/entity"
)

type UserStoreSqlite struct{}

func NewUserStoreSqlite() *UserStoreSqlite {
	return &UserStoreSqlite{}
}

func (u *UserStoreSqlite) Create(user *entity.User) error {
	fmt.Printf("Usuario: %q Criado no banco Sqlite", user.Name)
	return nil
}
