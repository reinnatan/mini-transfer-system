package main

import (
	"github.com/reinnatan/linkaja/delivery"
	"github.com/reinnatan/linkaja/repository"
	"github.com/reinnatan/linkaja/usecase"
)

func main() {
	accountRepository := repository.InitiateAccountBalanceRepository()
	accountUseCase := usecase.InitiateAccountBalanceUseCase(accountRepository)
	accountDelivery := delivery.InitiateAccountBalanceDelivery(accountUseCase)
	accountDelivery.Run()
}
