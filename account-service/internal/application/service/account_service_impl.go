package service

import (
	"account-service/internal/application/command"
	"account-service/internal/application/dto"
	"account-service/internal/application/query"
	"account-service/internal/domain/entity"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type AccountServiceImpl struct {
	AccountCommandHandler *command.AccountCommandHandler
	AccountQueryHandler   *query.AccountQueryHandler
}

func NewAccountServiceImpl(db *sql.DB) *AccountServiceImpl {
	return &AccountServiceImpl{
		AccountCommandHandler: command.NewAccountCommandHandler(db),
		AccountQueryHandler:   query.NewAccountQueryHandler(db),
	}
}

func (a *AccountServiceImpl) CreateAccount(command dto.CreateAccountCommand) (entity.Account, error) {

	account, err := a.AccountCommandHandler.CreateAccount(command)

	if err != nil {
		return account, err
	}

	return account, nil
}

func (a *AccountServiceImpl) Authenticate(auth dto.AuthenticationRequest) (dto.AuthenticationResponse, error) {

	account, err := a.AccountQueryHandler.GetAccountByEmail(auth.Email)

	if err != nil {
		return dto.AuthenticationResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password[:]), []byte(auth.Password[:]))

	if err != nil {
		return dto.AuthenticationResponse{}, err
	}

	return dto.AuthenticationResponse{Token: "success"}, nil
}
