package repomock

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/reinnatan/linkaja/domain"
	"github.com/reinnatan/linkaja/models"
)

type AccountAmountRepository struct {
	Conn *sql.DB
}

func InitiateAccountBalanceRepository() domain.AccountAmountRepository {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbPortParser, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatalf("Error parsing float")
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", dbHost, dbPortParser, dbUser, dbPassword, dbName)

	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	return &AccountAmountRepository{conn}
}

func (m *AccountAmountRepository) GetAccountDetail(accountNumber string) (models.AccountCustomer, error) {

	account := models.AccountCustomer{}
	if accountNumber == "500" {
		return account, fmt.Errorf("Account not found")
	}
	return account, nil
}

func (m *AccountAmountRepository) TransferAmount(ctx context.Context, fromAccountNumber string, requestBody models.RequestTransfer) (models.Response, error) {
	//simulation if from acccount sender not found
	if fromAccountNumber == "500" {
		response := models.Response{}
		response.ResponseCode = http.StatusBadRequest
		response.Message = "Account Sender not found"
		return response, fmt.Errorf(response.Message)
	}

	if requestBody.ToAccountNumber == "500" {
		response := models.Response{}
		response.ResponseCode = http.StatusBadRequest
		response.Message = "Account Receiver Not found"
		return response, fmt.Errorf(response.Message)
	}

	//simulation if balance is not enough
	if requestBody.Amount < 1000 {
		response := models.Response{}
		response.ResponseCode = http.StatusBadRequest
		response.Message = "Your Balance is not enough to transfer"
		return response, fmt.Errorf(response.Message)
	}

	response := models.Response{}
	response.ResponseCode = 201
	response.Message = "Successfully transfer amount"
	return response, nil
}
