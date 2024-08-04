package port

import "account-service/internal/domain/entity"

type AccountRepository interface {
	SaveAccount(account entity.Account) (entity.Account, error)
}
