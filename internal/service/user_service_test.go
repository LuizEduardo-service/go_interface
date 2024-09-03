package service_test

import (
	"log"
	"testing"

	"github.com/LuizEduardo-service/exe_interface/internal/entity"
	"github.com/LuizEduardo-service/exe_interface/internal/service"
)

type MockUserStore struct {
	users []*entity.User
}

func (m *MockUserStore) Create(user *entity.User) error {
	m.users = append(m.users, user)
	return nil
}

func TestCreateUser(t *testing.T) {
	userName := "Luiz Eduardo"

	mockUserStore := &MockUserStore{
		users: []*entity.User{},
	}

	userService, err := service.NewUserService(mockUserStore)
	if err != nil {
		log.Panicln(err)
	}

	userService.CreateUserHandler(userName)

	if mockUserStore.users[0] != nil {
		log.Fatalf("Usuario não cadastrado")
	}

}
