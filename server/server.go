package server

import (
	"github.com/engineering/CodeInformation/api"
	"github.com/engineering/CodeInformation/app"
	"github.com/gofiber/fiber/v2"
)

// Server is a interface to define contract to server up
type Server interface {
	Start()
}

type server struct {
	Fiber *fiber.App
	App   *app.Container
}

// New is instance the server
func New() Server {
	return &server{}
}

func (e *server) Start() {
	e.App = app.Register()
	e.Fiber = api.Register(e.App)
	e.Fiber.Listen(":3000")
}
