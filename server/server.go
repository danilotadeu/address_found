package server

import (
	"os"

	"github.com/danilotadeu/pismo/api"
	"github.com/danilotadeu/pismo/app"
	"github.com/danilotadeu/pismo/store"
	"github.com/gofiber/fiber/v2"
)

// Server is a interface to define contract to server up
type Server interface {
	Start()
}

type server struct {
	Fiber *fiber.App
	App   *app.Container
	Store *store.Container
}

// New is instance the server
func New() Server {
	return &server{}
}

func (e *server) Start() {
	e.Store = store.Register()
	e.App = app.Register(e.Store)
	e.Fiber = api.Register(e.App)
	e.Fiber.Listen(":" + os.Getenv("PORT"))
}
