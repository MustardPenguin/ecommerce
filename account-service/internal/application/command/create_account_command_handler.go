package command

import (
	"account-service/internal/application/dto"
	"account-service/internal/application/port"
	"account-service/internal/domain"
	"account-service/internal/domain/entity"
	"account-service/internal/infrastructure/repository/account"
	"database/sql"
)

type AccountCommandHandler struct {
	AccountDomainService *domain.AccountDomainService
	AccountRepository    port.AccountRepository
}

func NewAccountCommandHandler(db *sql.DB) *AccountCommandHandler {
	return &AccountCommandHandler{
		AccountDomainService: &domain.AccountDomainService{},
		AccountRepository:    account.NewAccountRepositoryImpl(db),
	}
}

func (a *AccountCommandHandler) CreateAccount(command dto.CreateAccountCommand) (entity.Account, error) {

	account := commandToAccount(command)
	err := a.AccountDomainService.ValidateCredentials(account)

	if err != nil {
		return entity.Account{}, err
	}

	account, err = a.AccountRepository.SaveAccount(account)

	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func commandToAccount(command dto.CreateAccountCommand) entity.Account {
	return entity.Account{
		Email:    command.Email,
		Password: command.Password,
	}
}
