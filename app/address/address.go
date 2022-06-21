package address

import (
	"context"

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
	response, err := a.integrations.ViaCep.Connect(ctx, zip)
	if err != nil {
		return nil, err
	}
	responseTransformed := address.TransformResult(response)
	return &responseTransformed, nil
}
