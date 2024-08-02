package api

import (
	"account-service/internal/api/controller"
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

	//accountServiceImpl := service.NewAccountServiceImpl()
	//
	//accountController := controller.AccountController{
	//	AccountService: accountServiceImpl,
	//}
	accountController := controller.NewAccountController()

	http.HandleFunc("POST /api/account", accountController.RegisterAccount)
}
