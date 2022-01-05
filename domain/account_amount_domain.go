package domain

import (
	"context"

	"github.com/reinnatan/linkaja/models"
)

type AccountAmountRepository interface {
	GetAccountDetail(accountNumber string) (models.AccountCustomer, error)
	TransferAmount(ctx context.Context, fromAccountNumber string, requestBody models.RequestTransfer) (models.Response, error)
}

type AccountBalanceUseCase interface {
	GetAccountDetail(accountNumber string) (models.AccountCustomer, error)
	TransferAmount(ctx context.Context, fromAccountNumber string, requestBody models.RequestTransfer) (models.Response, error)
}
