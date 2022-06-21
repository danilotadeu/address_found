package integrations

import (
	"log"

	"github.com/danilotadeu/address_found/integrations/viacep"
)

//Container ...
type Container struct {
	ViaCep viacep.Integrations
}

type Options struct {
	UrlViaCep string
}

//Register app container
func Register(options Options) *Container {
	container := &Container{
		ViaCep: viacep.NewViaCep(options.UrlViaCep),
	}
	log.Println("Registered -> Integrations")
	return container
}
