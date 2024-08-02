package repository

import (
	"account-service/internal/domain/entity"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"log"
	"testing"
)

const dbUser string = "user"
const dbPassword string = "admin"

var database *sql.DB

func setupDbContainer() *postgres.PostgresContainer {

	ctx := context.Background()

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

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	return postgresContainer
}

func connectToDb(pgContainer *postgres.PostgresContainer) {

	ctx := context.Background()
	dbHost, err := pgContainer.Host(ctx)

	if err != nil {
		log.Fatalf("error getting db host: %v", err)
	}

	dbPort := "8080"

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, "postgres")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing db")
		}
	}(db)

	database = db
}

func TestMain(m *testing.M) {
	pgContainer := setupDbContainer()
	connectToDb(pgContainer)
	m.Run()
}

func TestSaveAccount(t *testing.T) {

	accountRepositoryImpl := NewAccountRepositoryImpl(database)

	account := entity.Account{
		Email:    "test@test",
		Password: "password",
	}
	savedAccount, err := accountRepositoryImpl.SaveAccount(&account)

	if err != nil {
		t.Errorf("error saving account: %v", err)
	}
	if savedAccount.AccountId == 0 {
		t.Error("expected account id got 0")
	}
	if savedAccount.Email != account.Email {
		t.Errorf("expected %s got %s for email", savedAccount.Email, account.Email)
	}

}
