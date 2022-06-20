package app

import (
	"log"

	"github.com/danilotadeu/address_found/app/address"
)

//Container ...
type Container struct {
	Address address.App
}

type Options struct {
	UrlViaCep string
}

//Register app container
func Register(options Options) *Container {
	container := &Container{
		Address: address.NewApp(options.UrlViaCep),
	}
	log.Println("Registered -> App")
	return container
}
