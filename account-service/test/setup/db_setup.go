package setup

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"log"
)

var dbUser string = "user"
var dbPassword string = "password"

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

func Startup(ctx context.Context) (*postgres.PostgresContainer, *sql.DB) {

	pgContainer := setupDbContainer(ctx)
	db := connectToDb(pgContainer, ctx)

	ExecSql(db, "CREATE SCHEMA account;")
	ExecSql(db, "CREATE TABLE account.accounts (account_id SERIAL PRIMARY KEY, email VARCHAR(50) NOT NULL, password VARCHAR(255) NOT NULL);")

	return pgContainer, db
}

func ExecSql(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatalf("error setting up db: %v", err)
	}
}
