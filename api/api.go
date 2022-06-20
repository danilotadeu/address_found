package api

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/danilotadeu/address_found/api/address"
	"github.com/danilotadeu/address_found/app"
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

	// Address
	address.NewAPI(fiberRoute.Group("/address"), apps)

	log.Println("Registered -> Api")
	return fiberRoute
}
