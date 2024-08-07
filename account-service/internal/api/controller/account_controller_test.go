package controller

import (
	"account-service/internal/application/dto"
	"account-service/internal/domain/entity"
	"account-service/test/setup"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var database *sql.DB

func getBody[T any](t testing.TB, res *http.Response, data T) T {
	t.Helper()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Unable to read response body: %v", http.StatusInternalServerError)
	}
	defer res.Body.Close()

	err = json.Unmarshal(body, &data)
	if err != nil {
		t.Errorf("Error parsing json: %v", err)
	}

	return data
}

func convertToJson(t *testing.T, body map[string]interface{}) []byte {
	t.Helper()
	j, err := json.Marshal(body)

	if err != nil {
		t.Errorf("error while parsing json: %v", err)
	}

	return j
}

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
	body := map[string]interface{}{
		"email":    "test@test",
		"password": "test",
	}
	bodyJson := convertToJson(t, body)

	req, err := http.NewRequest("POST", "/api/account", bytes.NewBuffer(bodyJson))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	controller := NewAccountController(database)
	w := httptest.NewRecorder()
	controller.RegisterAccount(w, req)

	res := w.Result()
	got := getBody(t, res, entity.Account{})
	want := entity.Account{AccountId: 1, Email: "test@test"}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAuthenticate(t *testing.T) {
	body := map[string]interface{}{
		"email":    "test@test",
		"password": "test",
	}
	bodyJson := convertToJson(t, body)

	req, err := http.NewRequest("POST", "/api/authenticate", bytes.NewBuffer(bodyJson))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	controller := NewAccountController(database)
	w := httptest.NewRecorder()
	controller.Authenticate(w, req)

	res := w.Result()
	got := getBody(t, res, dto.AuthenticationResponse{})
	want := dto.AuthenticationResponse{Token: "success"}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
