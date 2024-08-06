package controller

import (
	"account-service/test/setup"
	"context"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var database *sql.DB

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

	database = db

	m.Run()
}

func TestCreateAccount(t *testing.T) {
	req, err := http.NewRequest("POST", "http://localhost:8081/api/account", nil)

	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	controller := NewAccountController(database)
	w := httptest.NewRecorder()
	controller.RegisterAccount(w, req)

	res := w.Result()
	log.Print(res)
}
