package main

import (
	serverInit "github.com/engineering/CodeInformation/server"
)

var (
	server serverInit.Server
)

func init() {
	server = serverInit.New()
}

func main() {
	server.Start()
}
