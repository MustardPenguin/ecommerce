package service

import (
	"account-service/internal/application/dto/command"
	"account-service/internal/domain"
	"account-service/internal/domain/entity"
	"fmt"
)

type AccountServiceImpl struct {
	AccountDomainService domain.AccountDomainService
}

func (a *AccountServiceImpl) CreateAccount(command command.CreateAccountCommand) entity.Account {
	fmt.Printf("called account service impl")
	account := commandToAccount(command)

	return account
}

func commandToAccount(command command.CreateAccountCommand) entity.Account {
	return entity.Account{
		Email:    command.Email,
		Password: command.Password,
	}
}
