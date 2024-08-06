package account

import (
	"account-service/internal/domain/entity"
	"account-service/test/setup"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"testing"
)

var accountRepos *AccountRepositoryImpl

func TestMain(m *testing.M) {

	ctx := context.Background()
	pgContainer, db := setup.Startup(ctx)

	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing db")
		}
	}(db)

	accountRepos = &AccountRepositoryImpl{DB: db}

	m.Run()
}

// UTIL
func saveAccount(t testing.TB, account entity.Account) entity.Account {
	t.Helper()
	savedAccount, err := accountRepos.SaveAccount(account)

	if err != nil {
		t.Errorf("error saving account: %v", err)
	}

	return savedAccount
}

// TESTS
func TestSaveAccount(t *testing.T) {
	account := entity.Account{
		AccountId: 1,
		Email:     "test@test",
		Password:  "password",
	}
	savedAccount := saveAccount(t, account)

	if savedAccount != account {
		t.Fatalf("expected %v got %v", account, savedAccount)
	}
}

func TestGetAccountByEmail(t *testing.T) {
	account := entity.Account{
		Email:    "this@mail",
		Password: "this",
	}

	savedAccount := saveAccount(t, account)

	got, err := accountRepos.GetAccountByEmail("this@mail")

	if err != nil {
		t.Errorf("error while getting email: %v", err)
	}

	if savedAccount != got {
		t.Errorf("got %v want %v", got, account)
	}
}

func TestGetAccountById(t *testing.T) {
	got, err := accountRepos.GetAccountById("1")

	if err != nil {
		t.Errorf("error while getting account by id: %v", err)
	}
	if got.AccountId != 1 {
		t.Errorf("expected account id 1 got %d", got.AccountId)
	}
	if got.Email != "test@test" {
		t.Errorf("expected email of test@test got %s", got.Email)
	}
}
