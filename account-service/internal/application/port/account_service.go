package port

import (
	"account-service/internal/application/dto/command"
	"account-service/internal/domain/entity"
)

type AccountService interface {
	CreateAccount(command command.CreateAccountCommand) entity.Account
}
