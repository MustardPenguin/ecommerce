package domain

import (
	"account-service/internal/domain/entity"
	"errors"
	"strings"
)

type AccountDomainService struct{}

func NewAccountDomainService() *AccountDomainService {
	return &AccountDomainService{}
}

func (a *AccountDomainService) ValidateCredentials(account entity.Account) error {
	err := a.ValidateEmail(account.Email)
	if err != nil {
		return err
	}
	return a.ValidatePassword(account.Password)
}

func (a *AccountDomainService) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("empty email string")
	}
	if len(email) > 50 {
		return errors.New("email above character limit")
	}
	if strings.Index(email, "@") == -1 {
		return errors.New("invalid email")
	}
	return nil
}

func (a *AccountDomainService) ValidatePassword(password string) error {
	length := len(password)
	if length < 3 {
		return errors.New("password length less than 3")
	}
	if length > 30 {
		return errors.New("password length greater than 30")
	}
	return nil
}
