package account

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

func (a *AccountRepositoryImpl) SaveAccount(account entity.Account) (entity.Account, error) {
	query := `INSERT INTO account.accounts (email, password) VALUES ($1, $2) RETURNING account_id`
	var id int64

	err := a.DB.QueryRow(query, &account.Email, &account.Password).Scan(&id)

	if err != nil {
		return entity.Account{}, err
	}

	account.AccountId = id
	return account, nil
}

func (a *AccountRepositoryImpl) GetAccountByEmail(email string) (entity.Account, error) {
	query := `SELECT account_id, email, password FROM account.accounts WHERE email = $1`
	return a.selectAccountQuery(query, email)
}

func (a *AccountRepositoryImpl) GetAccountById(id string) (entity.Account, error) {
	query := `SELECT account_id, email, password FROM account.accounts WHERE account_id = $1`
	return a.selectAccountQuery(query, id)
}

func (a *AccountRepositoryImpl) selectAccountQuery(query string, args ...string) (entity.Account, error) {
	var id int64
	var e, password string

	// Convert args from []string to []interface{}
	interfaceArgs := make([]interface{}, len(args))
	for i, v := range args {
		interfaceArgs[i] = v
	}

	err := a.DB.QueryRow(query, interfaceArgs...).Scan(&id, &e, &password)
	if err != nil {
		return entity.Account{}, err
	}
	return entity.Account{
		AccountId: id,
		Email:     e,
		Password:  password,
	}, nil
}
