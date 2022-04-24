package account

import (
	"context"
	"log"

	accountModel "github.com/danilotadeu/pismo/model/account"
	"github.com/danilotadeu/pismo/store"
)

//App is a contract to Account..
type App interface {
	CreateAccount(ctx context.Context, documentNumber string) (*int64, error)
	GetAccount(ctx context.Context, accountId int64) (*accountModel.AccountResultQuery, error)
}

type appImpl struct {
	store *store.Container
}

//NewApp init a account
func NewApp(store *store.Container) App {
	return &appImpl{
		store: store,
	}
}

//CreateAccount create a account..
func (a *appImpl) CreateAccount(ctx context.Context, documentNumber string) (*int64, error) {

	count, err := a.store.Account.GetAccountByDocumentNumber(ctx, documentNumber)
	if err != nil {
		log.Println("app.account.CreateAccount.GetAccountByDocumentNumber", err.Error())
		return nil, err
	}

	if count.Count > 0 {
		return nil, accountModel.ErrorAccountExists
	}

	id, err := a.store.Account.StoreCreateAccount(ctx, documentNumber)
	if err != nil {
		log.Println("app.account.CreateAccount.StoreCreateAccount", err.Error())
		return nil, err
	}
	return id, nil
}

//GetAccount get a account..
func (a *appImpl) GetAccount(ctx context.Context, accountId int64) (*accountModel.AccountResultQuery, error) {
	account, err := a.store.Account.GetAccount(ctx, accountId)
	if err != nil {
		log.Println("app.account.GetAccount.GetAccount", err.Error())
		return nil, err
	}
	return account, nil
}
