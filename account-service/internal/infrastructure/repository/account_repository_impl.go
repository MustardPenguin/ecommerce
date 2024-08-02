package repository

import (
	"account-service/internal/domain/entity"
	"database/sql"
)

type AccountRepositoryImpl struct {
	DB *sql.DB
}

func NewAccountRepositoryImpl(db *sql.DB) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{
		DB: db,
	}
}

func (a *AccountRepositoryImpl) SaveAccount(account *entity.Account) (*entity.Account, error) {
	query := `INSERT INTO account.accounts (email, password) VALUES ($1, $2) RETURNING id`

	err := a.DB.QueryRow(query, &account.Email, &account.Password).Scan(&account.AccountId)

	if err != nil {
		return nil, err
	}

	account.AccountId = 1
	return account, nil
}
