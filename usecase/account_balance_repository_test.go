package usecase_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/reinnatan/linkaja/models"
	repomock "github.com/reinnatan/linkaja/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountDetail(t *testing.T) {
	dbMock, _, err := sqlmock.New()

	repo := repomock.AccountAmountRepository{
		Conn: dbMock,
	}

	accountCustomer, err := repo.GetAccountDetail("10001")
	assert.NoError(t, err)
	assert.NotNil(t, accountCustomer)
}

func TestGetAccountDetailNotFound(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	repo := repomock.AccountAmountRepository{
		Conn: dbMock,
	}

	accountCustomer, err := repo.GetAccountDetail("500")
	assert.Error(t, err)
	assert.Equal(t, "", accountCustomer.AccountNumber)
}

func TestTransferAmmountSuccessullly(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub  database connection", err)
	}
	repo := repomock.AccountAmountRepository{
		Conn: dbMock,
	}
	body := models.RequestTransfer{}
	body.Amount = 10000
	body.ToAccountNumber = "550002"

	response, err := repo.TransferAmount(context.Background(), "234", body)
	assert.NoError(t, err)
	assert.Equal(t, 201, response.ResponseCode)
}

func TestTransferAmmountNotEnnough(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub  database connection", err)
	}
	repo := repomock.AccountAmountRepository{
		Conn: dbMock,
	}
	body := models.RequestTransfer{}
	body.Amount = 500
	body.ToAccountNumber = "550002"

	response, err := repo.TransferAmount(context.Background(), "234", body)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, response.ResponseCode)
}

func TestAccountSenderNotFound(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub  database connection", err)
	}
	repo := repomock.AccountAmountRepository{
		Conn: dbMock,
	}
	body := models.RequestTransfer{}
	body.Amount = 500
	body.ToAccountNumber = "550002"

	response, err := repo.TransferAmount(context.Background(), "500", body)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, response.ResponseCode)
}

func TestAccountReceiverNotFound(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub  database connection", err)
	}
	repo := repomock.AccountAmountRepository{
		Conn: dbMock,
	}
	body := models.RequestTransfer{}
	body.Amount = 1000
	body.ToAccountNumber = "500"

	response, err := repo.TransferAmount(context.Background(), "1000", body)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, response.ResponseCode)
}
