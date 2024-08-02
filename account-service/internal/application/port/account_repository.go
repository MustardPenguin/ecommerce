package port

import "account-service/internal/domain/entity"

type AccountRepository interface {
	SaveAccount() entity.Account
}
