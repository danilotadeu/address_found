package transaction

import "errors"

var ErrorTransactionTypeNotFound = errors.New("Transaction not fount")

var OperationTypes = map[int]string{1: "Compra a vista", 2: "Compra parcelada", 3: "Saque", 4: "Pagamento"}
var OperationTypesBuyOrWithdraw = map[int]string{1: "Compra a vista", 2: "Compra parcelada", 3: "Saque"}

//TransactionRequest is a struct to create account
type TransactionRequest struct {
	AccountID       int64   `json:"account_id" validate:"required"`
	OperationTypeID int     `json:"operation_type_id" validate:"required,min=1,max=4"`
	Amount          float64 `json:"amount" validate:"required,min=1"`
}

//TransactionResponse is a struct to response account
type TransactionResponse struct {
	ID int64 `json:"id"`
}
