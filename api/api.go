package api

import (
	"log"

	"github.com/danilotadeu/pismo/api/account"
	"github.com/danilotadeu/pismo/api/transaction"
	"github.com/danilotadeu/pismo/app"
	"github.com/gofiber/fiber/v2"
)

//Register ...
func Register(apps *app.Container) *fiber.App {
	fiberRoute := fiber.New()

	// Accounts
	account.NewAPI(fiberRoute.Group("/accounts"), apps)
	// Transactions
	transaction.NewAPI(fiberRoute.Group("/transactions"), apps)

	log.Println("Registered -> Api")
	return fiberRoute
}
