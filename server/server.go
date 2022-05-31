package server

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"

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
	Db    *sql.DB
}

// New is instance the server
func New() Server {
	return &server{}
}

func (e *server) Start() {
	e.Store, e.Db = store.Register()
	e.App = app.Register(e.Store)
	e.Fiber = api.Register(e.App)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = e.Fiber.Shutdown()
		_ = e.Db.Close()
	}()

	e.Fiber.Listen(":" + os.Getenv("PORT"))
}
