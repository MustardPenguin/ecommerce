package query

import "account-service/internal/application/port"

type AccountQueryHandler struct {
	AccountRepository port.AccountRepository
}