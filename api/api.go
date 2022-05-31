package api

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/danilotadeu/pismo/api/account"
	"github.com/danilotadeu/pismo/api/transaction"
	"github.com/danilotadeu/pismo/app"
	"github.com/gofiber/fiber/v2"
)

//Register ...
func Register(apps *app.Container) *fiber.App {
	fiberRoute := fiber.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = fiberRoute.Shutdown()
	}()

	// Accounts
	account.NewAPI(fiberRoute.Group("/accounts"), apps)
	// Transactions
	transaction.NewAPI(fiberRoute.Group("/transactions"), apps)

	log.Println("Registered -> Api")
	return fiberRoute
}
