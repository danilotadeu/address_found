package main

import (
	"os"

	serverInit "github.com/danilotadeu/r-customer-code-information/server"
)

var (
	server serverInit.Server
)

func init() {
	os.Setenv("URL_PROVIDER", "r-customer-code-information-provider-container")
	os.Setenv("PORT_PROVIDER", "4000")
	server = serverInit.New()
}

func main() {
	server.Start()
}
