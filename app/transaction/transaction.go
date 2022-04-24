package transaction

import (
	"context"
	"errors"
	"log"

	accountModel "github.com/danilotadeu/pismo/model/account"
	"github.com/danilotadeu/pismo/model/transaction"
	"github.com/danilotadeu/pismo/store"
)

//App is a contract to transaction..
type App interface {
	CreateTransaction(ctx context.Context, transactionBody transaction.TransactionRequest) (*int64, error)
}

type appImpl struct {
	store *store.Container
}

var GetAccount func(ctx context.Context, accountId int64) (*accountModel.AccountResultQuery, error)
var CreateTransaction func(ctx context.Context, accountID int64, operationTypeID int, amount float64) (*int64, error)

//NewApp init a transaction
func NewApp(store *store.Container) App {

	GetAccount = store.Account.GetAccount
	CreateTransaction = store.Transaction.CreateTransaction
	return &appImpl{
		store: store,
	}
}

//CreateAccount create a account..
func (a *appImpl) CreateTransaction(ctx context.Context, transactionBody transaction.TransactionRequest) (*int64, error) {
	_, ok := transaction.OperationTypes[transactionBody.OperationTypeID]
	if ok {
		_, err := GetAccount(ctx, transactionBody.AccountID)
		if err != nil {
			log.Println("app.transaction.CreateTransaction.GetAccount", err.Error())
			if errors.Is(err, accountModel.ErrorAccountNotFound) {
				return nil, accountModel.ErrorAccountNotFound
			}
			return nil, err
		}
		valueAmount := transactionBody.Amount
		_, ok := transaction.OperationTypesBuyOrWithdraw[transactionBody.OperationTypeID]
		if ok {
			valueAmount = -transactionBody.Amount
		}
		transactionID, err := CreateTransaction(ctx, transactionBody.AccountID, transactionBody.OperationTypeID, valueAmount)
		if err != nil {
			log.Println("app.transaction.CreateTransaction.CreateTransaction", err.Error())
			return nil, err
		}

		return transactionID, nil
	} else {
		return nil, transaction.ErrorTransactionTypeNotFound
	}
}
