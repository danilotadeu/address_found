package address

import (
	"context"
	"log"
	"net/http"

	"github.com/danilotadeu/address_found/app"
	addressModel "github.com/danilotadeu/address_found/model/address"
	errorsP "github.com/danilotadeu/address_found/model/errors_handler"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type apiImpl struct {
	apps *app.Container
}

// NewAPI address function..
func NewAPI(g fiber.Router, apps *app.Container) {
	api := apiImpl{
		apps: apps,
	}

	g.Post("/zip", api.zip)
}

func (p *apiImpl) zip(c *fiber.Ctx) error {
	bodyAddress := new(addressModel.ZipRequest)
	if err := c.BodyParser(bodyAddress); err != nil {
		log.Println("api.address.zip.body_parser", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errorsP.ErrorsResponse{
			Message: "Por favor enviar o cep como string.",
		})
	}

	validate := validator.New()
	if err := validate.Struct(bodyAddress); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errorsP.ErrorsResponse{
			Message: "O zip é obrigatório e necessita no mínimo 8 caracteres e no máximo 8 caracteres",
		})
	}

	ctx := context.Background()

	response, err := p.apps.Address.FindAddress(ctx, bodyAddress.Zip)
	if err != nil {
		log.Println("api.address.zip.FindAddress", err.Error())
		return c.Status(http.StatusInternalServerError).JSON(errorsP.ErrorsResponse{
			Message: "Por favor tente mais tarde...",
		})
	}

	return c.Status(http.StatusCreated).JSON(response)
}
