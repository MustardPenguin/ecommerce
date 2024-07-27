package domain

import (
	"account-service/internal/domain/entity"
	"testing"
)

func isErrorNil(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("Expected error but got none")
	}
}

func isError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error but got one: %v", err)
	}
}

func TestValidateInvalidEmail(t *testing.T) {
	accountDomainService := NewAccountDomainService()
	t.Run("Empty string", func(t *testing.T) {
		err := accountDomainService.ValidateEmail("")
		isErrorNil(t, err)
	})
	t.Run("Above max character limit", func(t *testing.T) {
		err := accountDomainService.ValidateEmail("asdjaslkdjaskldasdfsdfdsafadsfadsfdfadsfadsfjaklsjdlasjldaskdjaksjdaklsjdajsdakjdasd")
		isErrorNil(t, err)
	})
	t.Run("Not valid email", func(t *testing.T) {
		err := accountDomainService.ValidateEmail("hello")
		isErrorNil(t, err)
	})
}

func TestValidateValidEmail(t *testing.T) {
	accountDomainService := NewAccountDomainService()
	err := accountDomainService.ValidateEmail("test@test")
	isError(t, err)
}

func TestValidateInvalidPassword(t *testing.T) {
	accountDomainService := NewAccountDomainService()
	t.Run("Empty string / length less than 3", func(t *testing.T) {
		err := accountDomainService.ValidatePassword("")
		isErrorNil(t, err)
	})
	t.Run("Password greater than 30", func(t *testing.T) {
		err := accountDomainService.ValidateEmail("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		isErrorNil(t, err)
	})
}

func TestValidateValidPassword(t *testing.T) {
	accountDomainService := NewAccountDomainService()
	err := accountDomainService.ValidatePassword("password")
	isError(t, err)
}

func TestValidateValidCredentials(t *testing.T) {
	accountDomainService := NewAccountDomainService()
	err := accountDomainService.ValidateCredentials(entity.Account{AccountId: 1, Email: "sadfsd@fdas", Password: "adsfasdfasd"})
	isError(t, err)
}

func TestValidateInvalidCredentials(t *testing.T) {
	t.Run("invalid email", func(t *testing.T) {
		accountDomainService := NewAccountDomainService()
		err := accountDomainService.ValidateCredentials(entity.Account{AccountId: 1, Email: "aaa", Password: "asdasd"})
		isErrorNil(t, err)
	})
	t.Run("invalid password", func(t *testing.T) {
		accountDomainService := NewAccountDomainService()
		err := accountDomainService.ValidateCredentials(entity.Account{AccountId: 1, Email: "aaa@aaa", Password: ""})
		isErrorNil(t, err)
	})
}
