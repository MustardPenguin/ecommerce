package api

import (
	"account-service/internal/api/controller"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Server struct{}

func StartServer(db *sql.DB, port string) {
	setupController(db)

	addr := fmt.Sprintf(":%s", port)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalf("Error running http server: %v", err)
	}
}

func setupController(db *sql.DB) {

	accountController := controller.NewAccountController(db)

	http.HandleFunc("POST /api/account", accountController.RegisterAccount)
	http.HandleFunc("POST /api/authenticate", accountController.Authenticate)
}
