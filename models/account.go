package models

type AccountCustomer struct {
	AccountNumber string `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int64  `json:"balance"`
}

type Response struct {
	ResponseCode int    `json:"response"`
	Message      string `json:"messsage"`
}

type RequestTransfer struct {
	ToAccountNumber string `json:"to_account_number"`
	Amount          int64  `json"amount"`
}
