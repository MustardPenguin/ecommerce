package controller

import (
	"account-service/internal/api/helper"
	"account-service/internal/application/dto"
	"account-service/internal/application/port"
	"account-service/internal/application/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AccountController struct {
	AccountService port.AccountService
}

func NewAccountController(db *sql.DB) *AccountController {
	return &AccountController{
		AccountService: service.NewAccountServiceImpl(db),
	}
}

func (a *AccountController) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	command := helper.GetBody(w, r, dto.CreateAccountCommand{})
	log.Print(command)
	account, err := a.AccountService.CreateAccount(command)

	if err != nil {
		errStr := fmt.Sprintf("Error creating account: %v", err)
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		errStr := fmt.Sprintf("Error encoding json: %v", err)
		log.Print(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
	}
}

func (a *AccountController) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := helper.GetBody(w, r, dto.AuthenticationRequest{})
	res, err := a.AccountService.Authenticate(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		errStr := fmt.Sprintf("Error encoding json: %v", err)
		log.Print(errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
	}
}
