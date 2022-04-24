package account

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/danilotadeu/pismo/app"
	accountModel "github.com/danilotadeu/pismo/model/account"
	errorsP "github.com/danilotadeu/pismo/model/errors_handler"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type apiImpl struct {
	apps *app.Container
}

// NewAPI account function..
func NewAPI(g fiber.Router, apps *app.Container) {
	api := apiImpl{
		apps: apps,
	}

	g.Post("/", api.accountCreate)
	g.Get("/:accountId", api.account)
}

func (p *apiImpl) accountCreate(c *fiber.Ctx) error {
	bodyAccount := new(accountModel.AccountRequest)
	if err := c.BodyParser(bodyAccount); err != nil {
		log.Println("api.account.accountCreate.body_parser", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errorsP.ErrorsResponse{
			Message: "Por favor tente mais tarde...",
		})
	}

	validate := validator.New()
	if err := validate.Struct(bodyAccount); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errorsP.ErrorsResponse{
			Message: err.Error(),
		})
	}

	ctx := context.Background()

	response, err := p.apps.Account.CreateAccount(ctx, bodyAccount.DocumentNumber)
	if err != nil {
		log.Println("api.account.accountCreate.CreateAccount", err.Error())
		if errors.Is(err, accountModel.ErrorAccountExists) {
			return c.Status(http.StatusNotFound).JSON(errorsP.ErrorsResponse{
				Message: fmt.Sprintf("Já existe uma conta com esse documento (%s)", bodyAccount.DocumentNumber),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(errorsP.ErrorsResponse{
			Message: "Por favor tente mais tarde...",
		})
	}

	return c.Status(http.StatusCreated).JSON(accountModel.AccountResponse{
		ID: *response,
	})
}

func (p *apiImpl) account(c *fiber.Ctx) error {
	accountId := c.Params("accountId")
	if len(accountId) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(
			errorsP.ErrorsResponse{
				Message: "Por favor informe o accountId",
			},
		)
	}

	iAccountId, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		log.Println("api.account.account.ParseInt", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errorsP.ErrorsResponse{
			Message: "Por favor tente mais tarde...",
		})
	}

	ctx := context.Background()
	account, err := p.apps.Account.GetAccount(ctx, iAccountId)
	if err != nil {
		log.Println("api.account.account.GetAccount", err.Error())
		if errors.Is(err, accountModel.ErrorAccountNotFound) {
			return c.Status(http.StatusNotFound).JSON(errorsP.ErrorsResponse{
				Message: fmt.Sprintf("Conta (%d) não encontrada", iAccountId),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(errorsP.ErrorsResponse{
			Message: "Por favor tente mais tarde...",
		})
	}

	return c.Status(http.StatusOK).JSON(account)
}
