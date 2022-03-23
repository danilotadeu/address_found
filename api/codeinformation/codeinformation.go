package codeinformation

import (
	"github.com/engineering/CodeInformation/app"
	codeinformationModel "github.com/engineering/CodeInformation/model/codeinformation"
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

	g.Post("/", api.CodeInformation)
}

func (p *apiImpl) CodeInformation(c *fiber.Ctx) error {
	requestCodeInformation := new(codeinformationModel.CodeInformationRequest)
	if err := c.BodyParser(requestCodeInformation); err != nil {
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
	customerCode := p.apps.CodeInformation.GetCodeInformation(ctx, requestCodeInformation, clientID, messageID)

	return c.JSON(codeinformationModel.CodeInformationResponse{
		CustomerCode: customerCode,
	})
}
