package codeinformation

import (
	"log"

	"github.com/danilotadeu/r-customer-code-information/app"
	codeinformationModel "github.com/danilotadeu/r-customer-code-information/model/codeinformation"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type apiImpl struct {
	apps *app.Container
}

// NewAPI codeinformation function..
func NewAPI(g fiber.Router, apps *app.Container) {
	api := apiImpl{
		apps: apps,
	}

	g.Post("/", api.CodeInformationHandler)
}

func (p *apiImpl) CodeInformationHandler(c *fiber.Ctx) error {
	requestCodeInformation := new(codeinformationModel.CodeInformationRequest)
	if err := c.BodyParser(requestCodeInformation); err != nil {
		log.Println("api.codeinformation.codeinformation.codeinformation.body_parser", err.Error())
		return err
	}

	headers := c.GetReqHeaders()
	var clientID string
	if val, ok := headers["clientId"]; ok {
		clientID = val
	} else {
		clientID = uuid.New().String()
	}

	var messageID string
	if val, ok := headers["messageId"]; ok {
		messageID = val
	} else {
		messageID = uuid.New().String()
	}

	ctx := c.Context()
	customerCode, err := p.apps.CodeInformation.GetCodeInformation(ctx, requestCodeInformation, clientID, messageID)
	if err != nil {
		log.Println("api.codeinformation.codeinformation.codeinformation.get_code_information", err.Error())
		return c.JSON(codeinformationModel.CodeInformationResponseError{
			Description: err.Error(),
			Provider: codeinformationModel.Provider{
				ServiceName:  "r-customer-code-information",
				ErrorCode:    "500",
				ErrorMessage: err.Error(),
			},
		})
	}

	return c.JSON(codeinformationModel.CodeInformationResponse{
		CustomerCode: *customerCode,
	})
}