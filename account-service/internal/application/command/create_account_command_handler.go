package command

import (
	"account-service/internal/application/dto"
	"account-service/internal/domain/entity"
)

type AccountCommandHandler struct{}

func (a *AccountCommandHandler) CreateAccount(command dto.CreateAccountCommand) entity.Account {

	account := commandToAccount(command)

	return account
}

func commandToAccount(command dto.CreateAccountCommand) entity.Account {
	return entity.Account{
		Email:    command.Email,
		Password: command.Password,
	}
}
