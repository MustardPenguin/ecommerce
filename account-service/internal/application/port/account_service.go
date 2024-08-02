package port

import (
	"account-service/internal/application/dto"
	"account-service/internal/domain/entity"
)

type AccountService interface {
	CreateAccount(command dto.CreateAccountCommand) entity.Account
}
