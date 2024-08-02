package service

import (
	"account-service/internal/application/command"
	"account-service/internal/application/dto"
	"account-service/internal/domain"
	"account-service/internal/domain/entity"
	"fmt"
)

type AccountServiceImpl struct {
	AccountDomainService  domain.AccountDomainService
	AccountCommandHandler command.AccountCommandHandler
}

func NewAccountServiceImpl() *AccountServiceImpl {
	return &AccountServiceImpl{
		AccountDomainService:  domain.AccountDomainService{},
		AccountCommandHandler: command.AccountCommandHandler{},
	}
}

func (a *AccountServiceImpl) CreateAccount(command dto.CreateAccountCommand) entity.Account {
	fmt.Printf("called account service impl")

	return a.AccountCommandHandler.CreateAccount(command)
}
