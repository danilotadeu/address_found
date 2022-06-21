package address

import (
	"context"
	"log"

	"github.com/danilotadeu/address_found/integrations"
	"github.com/danilotadeu/address_found/model/address"
)

//App is a contract to Address..
type App interface {
	FindAddress(ctx context.Context, zip string) (*address.ZipResponse, error)
}

type appImpl struct {
	integrations integrations.Container
}

//NewApp init a Address
func NewApp(integrations *integrations.Container) App {
	return &appImpl{
		integrations: *integrations,
	}
}

//FindAddress find a address in viacep or in another integrations..
func (a *appImpl) FindAddress(ctx context.Context, zip string) (*address.ZipResponse, error) {
	cResponseViaCep := make(chan *address.ResponseViaCep)

	go func() error {
		response, err := a.integrations.ViaCep.Connect(ctx, zip)
		if err != nil {
			log.Println("app.address.FindAddress.ViaCep.Connect", err.Error())
			return err
		}
		cResponseViaCep <- response
		return nil
	}()

	cResponseApiCep := make(chan *address.ResponseApiCep)

	go func() error {
		response, err := a.integrations.ApiCep.Connect(ctx, zip)
		if err != nil {
			log.Println("app.address.FindAddress.ApiCep.Connect", err.Error())
			return err
		}
		cResponseApiCep <- response
		return nil
	}()

	for i := 0; i < 2; i++ {
		select {
		case responseViaCepMsg := <-cResponseViaCep:
			responseTransformed := address.TransformResultViaCep(responseViaCepMsg)
			return &responseTransformed, nil

		case responseApiCepMsg := <-cResponseApiCep:
			responseTransformed := address.TransformResultApiCep(responseApiCepMsg)
			return &responseTransformed, nil
		}
	}

	return &address.ZipResponse{}, nil
}
