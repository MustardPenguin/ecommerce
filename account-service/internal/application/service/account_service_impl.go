package service

import (
	"account-service/internal/application/command"
	"account-service/internal/application/dto"
	"account-service/internal/domain/entity"
	"database/sql"
)

type AccountServiceImpl struct {
	AccountCommandHandler *command.AccountCommandHandler
}

func NewAccountServiceImpl(db *sql.DB) *AccountServiceImpl {
	return &AccountServiceImpl{
		AccountCommandHandler: command.NewAccountCommandHandler(db),
	}
}

func (a *AccountServiceImpl) CreateAccount(command dto.CreateAccountCommand) (entity.Account, error) {

	account, err := a.AccountCommandHandler.CreateAccount(command)

	if err != nil {
		return account, err
	}

	return account, nil
}
