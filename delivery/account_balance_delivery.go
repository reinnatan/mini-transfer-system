package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/reinnatan/linkaja/domain"
	"github.com/reinnatan/linkaja/models"
)

type AccountBalanceDelivery struct {
	accountBalanceUseCase domain.AccountBalanceUseCase
	router                *mux.Router
}

func InitiateAccountBalanceDelivery(accountUseCase domain.AccountBalanceUseCase) AccountBalanceDelivery {
	handler := AccountBalanceDelivery{
		accountBalanceUseCase: accountUseCase,
	}

	handler.router = mux.NewRouter()
	handler.router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("Hellooo World..."))
	}).Methods("GET")
	handler.router.HandleFunc("/account/{account_number}", handler.GetAccountDetail).Methods("GET")
	handler.router.HandleFunc("/account/{from_account_number}/transfer", handler.TransferAmount).Methods("POST")
	return handler
}

func (a *AccountBalanceDelivery) GetAccountDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	accountNumber := vars["account_number"]
	accountModel, error := a.accountBalanceUseCase.GetAccountDetail(accountNumber)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := models.Response{}
		response.ResponseCode = 400
		response.Message = "Account not found"
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(accountModel)
}

func (a *AccountBalanceDelivery) TransferAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	accountNumber := vars["from_account_number"]
	ctx := context.Background()

	requestBody := models.RequestTransfer{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		response := models.Response{}
		response.ResponseCode = 400
		response.Message = "Failed Parsing body"
		json.NewEncoder(w).Encode(response)
		return
	}

	response, _ := a.accountBalanceUseCase.TransferAmount(ctx, accountNumber, requestBody)
	json.NewEncoder(w).Encode(response)
}

func (a *AccountBalanceDelivery) Run() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	appsURL := os.Getenv("APPSURL")
	log.Fatal(http.ListenAndServe(appsURL, a.router))
}
