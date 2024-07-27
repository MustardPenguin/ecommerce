package api

import (
	"account-service/internal/api/controller"
	"account-service/internal/application/service"
	"account-service/internal/domain"
	"fmt"
	"log"
	"net/http"
)

type Server struct{}

func StartServer(port string) {
	setupController()

	addr := fmt.Sprintf(":%s", port)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalf("Error running http server: %v", err)
	}
}

func setupController() {
	accountServiceImpl := service.AccountServiceImpl{
		AccountDomainService: domain.AccountDomainService{},
	}

	accountController := controller.AccountController{
		AccountService: &accountServiceImpl,
	}

	http.HandleFunc("POST /api/account", accountController.RegisterAccount)
}
