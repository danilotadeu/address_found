package app

import (
	"log"

	"github.com/danilotadeu/address_found/app/address"
	"github.com/danilotadeu/address_found/integrations"
)

//Container ...
type Container struct {
	Address address.App
}

type Options struct {
	UrlViaCep string
}

//Register app container
func Register(integrations *integrations.Container) *Container {
	container := &Container{
		Address: address.NewApp(integrations),
	}
	log.Println("Registered -> App")
	return container
}
