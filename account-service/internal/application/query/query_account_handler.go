package query

import (
	"account-service/internal/application/port"
	"account-service/internal/domain/entity"
	"account-service/internal/infrastructure/repository/account"
	"database/sql"
)

type AccountQueryHandler struct {
	AccountRepository port.AccountRepository
}

func NewAccountQueryHandler(db *sql.DB) *AccountQueryHandler {
	return &AccountQueryHandler{
		AccountRepository: account.NewAccountRepositoryImpl(db),
	}
}

func (a *AccountQueryHandler) GetAccountByEmail(email string) (entity.Account, error) {
	account, err := a.AccountRepository.GetAccountByEmail(email)

	if err != nil {
		return entity.Account{}, nil
	}

	return account, nil
}
