package controller

import (
	"account-service/internal/application/dto/command"
	"account-service/internal/application/port"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AccountController struct {
	AccountService port.AccountService
}

func (a *AccountController) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	account := a.AccountService.CreateAccount(command.CreateAccountCommand{})

	err := json.NewEncoder(w).Encode(account)
	if err != nil {
		errStr := fmt.Sprintf("Error creating account: %v", err)
		log.Print(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
	}
}
