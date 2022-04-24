package account

import "errors"

var ErrorAccountNotFound = errors.New("Account not found")
var ErrorAccountExists = errors.New("Account yet exists")

//AccountRequest is a struct to create account
type AccountRequest struct {
	DocumentNumber string `json:"document_number" validate:"required,min=11"`
}

//AccountResponse is a struct to response account
type AccountResponse struct {
	ID int64 `json:"id"`
}

//AccountResultQuery is a struct to result query account
type AccountResultQuery struct {
	ID             int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

//AccountCountResultQuery is a struct to result query account count
type AccountCountResultQuery struct {
	Count int64 `json:"count"`
}
