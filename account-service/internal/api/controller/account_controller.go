package controller

import (
	"account-service/internal/application/dto"
	"account-service/internal/application/port"
	"account-service/internal/application/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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

	command := getBody(w, r, dto.CreateAccountCommand{})
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

func getBody[T any](w http.ResponseWriter, r *http.Request, data T) T {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return data
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "Error closing request body", http.StatusInternalServerError)
			return
		}
	}(r.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return data
	}

	return data
}
