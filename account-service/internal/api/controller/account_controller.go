package controller

import (
	"account-service/internal/application/dto"
	"account-service/internal/application/port"
	"account-service/internal/application/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AccountController struct {
	AccountService port.AccountService
}

func NewAccountController() *AccountController {
	return &AccountController{
		AccountService: service.NewAccountServiceImpl(),
	}
}

func (a *AccountController) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	account := a.AccountService.CreateAccount(dto.CreateAccountCommand{})

	err := json.NewEncoder(w).Encode(account)
	if err != nil {
		errStr := fmt.Sprintf("Error creating account: %v", err)
		log.Print(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
	}
}
