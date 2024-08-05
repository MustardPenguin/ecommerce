package command

import (
	"account-service/internal/application/dto"
	"account-service/internal/application/port"
	"account-service/internal/domain"
	"account-service/internal/domain/entity"
	"account-service/internal/infrastructure/repository/account"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
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

	found, _ := a.AccountRepository.GetAccountByEmail(command.Email)

	if found != (entity.Account{}) {
		return entity.Account{}, errors.New("email already taken")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(account.Password), 10)

	if err != nil {
		return entity.Account{}, err
	}

	account.Password = string(pw[:])
	account, err = a.AccountRepository.SaveAccount(account)

	if err != nil {
		return entity.Account{}, err
	}

	account.Password = ""
	return account, nil
}

func commandToAccount(command dto.CreateAccountCommand) entity.Account {
	return entity.Account{
		Email:    command.Email,
		Password: command.Password,
	}
}
