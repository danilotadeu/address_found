package integrations

import (
	"log"

	"github.com/danilotadeu/address_found/integrations/apicep"
	"github.com/danilotadeu/address_found/integrations/viacep"
)

//Container ...
type Container struct {
	ViaCep viacep.Integrations
	ApiCep apicep.Integrations
}

type Options struct {
	UrlViaCep string
	UrlApiCep string
}

//Register app container
func Register(options Options) *Container {
	container := &Container{
		ViaCep: viacep.NewViaCep(options.UrlViaCep),
		ApiCep: apicep.NewApiCep(options.UrlApiCep),
	}
	log.Println("Registered -> Integrations")
	return container
}
