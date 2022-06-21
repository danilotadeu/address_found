package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/danilotadeu/address_found/api"
	"github.com/danilotadeu/address_found/app"
	"github.com/danilotadeu/address_found/integrations"
	"github.com/gofiber/fiber/v2"
)

// Server is a interface to define contract to server up
type Server interface {
	Start()
}

type server struct {
	Fiber        *fiber.App
	App          *app.Container
	Integrations *integrations.Container
}

// New is instance the server
func New() Server {
	return &server{}
}

func (e *server) Start() {
	e.Integrations = integrations.Register(integrations.Options{
		UrlViaCep: os.Getenv("URL_VIACEP"),
		UrlApiCep: os.Getenv("URL_APICEP"),
	})
	e.App = app.Register(e.Integrations)
	e.Fiber = api.Register(e.App)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = e.Fiber.Shutdown()
	}()

	e.Fiber.Listen(":" + os.Getenv("PORT"))
}
