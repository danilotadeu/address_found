package api

import (
	"log"

	"github.com/danilotadeu/r-customer-code-information/api/codeinformation"
	"github.com/danilotadeu/r-customer-code-information/app"
	"github.com/gofiber/fiber/v2"
)

//Register ...
func Register(apps *app.Container) *fiber.App {
	fiberRoute := fiber.New()
	//Register group v1
	api := fiberRoute.Group("/v1")

	//Register group routes
	codeinformation.NewAPI(api.Group("/r-customer-code-information"), apps)

	log.Println("Registered -> Api")
	return fiberRoute
}
