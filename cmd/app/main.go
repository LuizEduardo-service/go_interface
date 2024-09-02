package main

import (
	"fmt"

	"github.com/LuizEduardo-service/exe_interface/internal/infra"
	"github.com/LuizEduardo-service/exe_interface/internal/service"
)

func main() {

	store := infra.NewUserStorePostgres()
	service, err := service.NewUserService(store)

	if err != nil {
		fmt.Println("Erro ao Criar Serviço")
	}

	service.CreateUserHandler("Luiz Eduardo")
}
