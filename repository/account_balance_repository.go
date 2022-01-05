package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

var temp *sql.DB

func InitiateAccountBalanceRepository() domain.AccountAmountRepository {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT") //"5432"

	dbPortParser, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatalf("Error parsing float")
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", dbHost, dbPortParser, dbUser, dbPassword, dbName)

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Printf("Terjadi Error Database %s\n", err.Error())
		log.Fatal(err)
	}

	temp = conn
	fmt.Printf("Berhasil Connect Database %s\n", connectionString)
	return &AccountAmountRepository{conn}
}

func (m *AccountAmountRepository) GetAccountDetail(accountNumber string) (models.AccountCustomer, error) {

	err := temp.Ping()
	if err != nil {
		fmt.Printf("Terjadi Error account detail %s\n", err.Error())
		os.Exit(1)
	}

	account := models.AccountCustomer{}
	statement := "SELECT c.customer_number, c.name, a.balance  FROM Customers c INNER JOIN Account a ON c.customer_number=a.customer_number where a.customer_number='" + accountNumber + "'"
	error := temp.QueryRow(statement).Scan(&account.AccountNumber, &account.CustomerName, &account.Balance)
	if error != nil {
		fmt.Printf("Terjadi Error %s", error.Error())
		return account, error
	}
	fmt.Println("Succes Get Account Detail")
	return account, nil
}

func (m *AccountAmountRepository) TransferAmount(ctx context.Context, fromAccountNumber string, requestBody models.RequestTransfer) (models.Response, error) {
	accountFrom := models.AccountCustomer{}
	accountFromQuery := "SELECT c.customer_number, c.name, a.balance  FROM Customers c INNER JOIN Account a ON c.customer_number=a.customer_number where a.account_number='" + fromAccountNumber + "'"
	error := m.Conn.QueryRow(accountFromQuery).Scan(&accountFrom.AccountNumber, &accountFrom.CustomerName, &accountFrom.Balance)
	if error != nil {
		response := models.Response{}
		response.ResponseCode = 400
		response.Message = "Account Sender not found"
		return response, error
	}

	accountTo := models.AccountCustomer{}
	accountToQuery := "SELECT c.customer_number, c.name, a.balance  FROM Customers c INNER JOIN Account a ON c.customer_number=a.customer_number where a.account_number='" + requestBody.ToAccountNumber + "'"
	error = m.Conn.QueryRow(accountToQuery).Scan(&accountTo.AccountNumber, &accountTo.CustomerName, &accountTo.Balance)
	if error != nil {
		response := models.Response{}
		response.ResponseCode = 400
		response.Message = "Account Receiver not found"
		return response, error
	}

	if accountFrom.Balance-requestBody.Amount < 0 {
		response := models.Response{}
		response.ResponseCode = 400
		response.Message = "Your Balance is not enough to transfer"
		return response, fmt.Errorf(response.Message)
	}

	tx, err := m.Conn.BeginTx(ctx, nil)
	if err != nil {
		response := models.Response{}
		response.ResponseCode = 400
		response.Message = "Something wrong when create transaction"
		return response, fmt.Errorf(response.Message)
	}

	// update balance from sender
	sqlUpdate := fmt.Sprintf("Update account set balance=%d where account_number='%s'", accountFrom.Balance-requestBody.Amount, fromAccountNumber)
	_, err = tx.ExecContext(ctx, sqlUpdate)
	if err != nil {
		tx.Rollback()
		response := models.Response{}
		response.ResponseCode = 400
		response.Message = "Something wrong when doing transfer"
		return response, fmt.Errorf(response.Message)
	}

	// update balance to reciver
	_, err = tx.ExecContext(ctx, fmt.Sprintf("Update account set balance=%d where account_number='%s'", accountTo.Balance+requestBody.Amount, requestBody.ToAccountNumber))
	if err != nil {
		tx.Rollback()
		response := models.Response{}
		response.ResponseCode = 400
		response.Message = "Something wrong when doing transfer"
		return response, fmt.Errorf(response.Message)
	}
	tx.Commit()

	response := models.Response{}
	response.ResponseCode = 201
	response.Message = "Successfully transfer amount"
	return response, nil
}
