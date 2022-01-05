package usecasemock

import (
	"context"

	"github.com/reinnatan/linkaja/domain"
	"github.com/reinnatan/linkaja/models"
	repomock "github.com/reinnatan/linkaja/repository/mock"
)

type AccountBalanceUseCase struct {
	accountBalanceRepo repomock.AccountAmountRepository
}

func InitiateAccountBalanceUseCase(domain repomock.AccountAmountRepository) domain.AccountBalanceUseCase {
	return &AccountBalanceUseCase{
		accountBalanceRepo: domain,
	}
}

func (a *AccountBalanceUseCase) GetAccountDetail(accountNumber string) (models.AccountCustomer, error) {
	return a.accountBalanceRepo.GetAccountDetail(accountNumber)
}

func (a *AccountBalanceUseCase) TransferAmount(ctx context.Context, fromAccountNumber string, requestBody models.RequestTransfer) (models.Response, error) {
	return a.accountBalanceRepo.TransferAmount(ctx, fromAccountNumber, requestBody)
}
