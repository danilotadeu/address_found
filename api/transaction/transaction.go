package transaction

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/danilotadeu/pismo/app"
	accountModel "github.com/danilotadeu/pismo/model/account"
	errorsP "github.com/danilotadeu/pismo/model/errors_handler"
	transactionModel "github.com/danilotadeu/pismo/model/transaction"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type apiImpl struct {
	apps *app.Container
}

var createTransactionFunc func(ctx context.Context, transactionBody transactionModel.TransactionRequest) (*int64, error)

// NewAPI account function..
func NewAPI(g fiber.Router, apps *app.Container) {
	api := apiImpl{
		apps: apps,
	}

	createTransactionFunc = apps.Transaction.CreateTransaction
	g.Post("/", api.transactionCreate)
}

func (p *apiImpl) transactionCreate(c *fiber.Ctx) error {
	bodyTransaction := new(transactionModel.TransactionRequest)
	if err := c.BodyParser(bodyTransaction); err != nil {
		log.Println("api.transaction.transactionCreate.body_parser", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errorsP.ErrorsResponse{
			Message: "Por favor tente mais tarde...",
		})
	}

	validate := validator.New()
	if err := validate.Struct(bodyTransaction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errorsP.ErrorsResponse{
			Message: err.Error(),
		})
	}

	ctx := context.Background()

	transactionID, err := createTransactionFunc(ctx, *bodyTransaction)
	if err != nil {
		log.Println("api.transaction.transactionCreate.CreateTransaction", err.Error())
		if errors.Is(err, accountModel.ErrorAccountNotFound) {
			return c.Status(http.StatusNotFound).JSON(errorsP.ErrorsResponse{
				Message: fmt.Sprintf("Conta (%d) não encontrada", bodyTransaction.AccountID),
			})
		}
		if errors.Is(err, transactionModel.ErrorTransactionTypeNotFound) {
			return c.Status(http.StatusNotFound).JSON(errorsP.ErrorsResponse{
				Message: fmt.Sprintf("Operação (%d) não encontrada", bodyTransaction.OperationTypeID),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(errorsP.ErrorsResponse{
			Message: "Por favor tente mais tarde...",
		})
	}

	return c.Status(http.StatusCreated).JSON(transactionModel.TransactionResponse{
		ID: *transactionID,
	})
}
