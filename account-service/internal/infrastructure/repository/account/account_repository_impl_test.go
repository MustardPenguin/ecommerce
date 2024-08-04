package account

import (
	"account-service/internal/domain/entity"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"log"
	"testing"
)

const dbUser string = "user"
const dbPassword string = "admin"

var accountRepos *AccountRepositoryImpl

// SETUP //
func setupDbContainer(ctx context.Context) *postgres.PostgresContainer {

	postgresContainer, err := postgres.Run(
		ctx,
		"postgres:latest",
		postgres.WithDatabase("postgres"),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)

	if err != nil {
		log.Fatalf("Error starting container: %v", err)
	}

	return postgresContainer
}

func connectToDb(pgContainer *postgres.PostgresContainer, ctx context.Context) *sql.DB {

	_, err := pgContainer.Host(ctx)

	if err != nil {
		log.Fatalf("error getting db host: %v", err)
	}

	connectionString, err := pgContainer.ConnectionString(ctx, "sslmode=disable")

	if err != nil {
		log.Fatalf("error getting connection string: %v", err)
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("error connecting do database: %v", err)
	}
	return db
}

func setupDb(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatalf("error setting up db: %v", err)
	}
}

func TestMain(m *testing.M) {

	ctx := context.Background()
	pgContainer := setupDbContainer(ctx)
	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()
	db := connectToDb(pgContainer, ctx)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing db")
		}
	}(db)

	setupDb(db, "CREATE SCHEMA account;")
	setupDb(db, "CREATE TABLE account.accounts (account_id SERIAL PRIMARY KEY, email VARCHAR(50) NOT NULL, password VARCHAR(255) NOT NULL);")

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
