package codeinformation

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danilotadeu/r-customer-code-information/app"
	codeinformationModel "github.com/danilotadeu/r-customer-code-information/model/codeinformation"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	messageIDC = "Messageid"
	clientIDC  = "clientId"
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
	if val, ok := headers[clientIDC]; ok {
		clientID = val
	}

	var messageID string
	if val, ok := headers[messageIDC]; ok {
		messageID = val
	} else {
		messageID = uuid.New().String()
	}
	c.Set(messageIDC, messageID)

	ctx := context.Background()
	ctx = context.WithValue(ctx, clientIDC, clientID)
	ctx = context.WithValue(ctx, messageIDC, messageID)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	/*LOG INICIO*/
	now := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	log.Println("Log one startup Level [debug]")
	log.Println("    INFO    EngDB | Engineering | GMID |  | - | ENGNB003153 | orch-r-customer-code-information | 1 | - | - | " + formatted + " | - | - | requestStarting | HTTP ID []")
	headersJSON, _ := json.Marshal(headers)
	log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | orch-r-customer-code-information | 1 | - | - | " + formatted + " | - | - | requestHeaderReceived | " + string(headersJSON))
	payload, _ := json.Marshal(requestCodeInformation)
	log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | orch-r-customer-code-information | 1 | - | - | " + formatted + " | - | - | requestPayloadReceived | " + string(payload))
	log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | orch-r-customer-code-information | 1 | - | - | " + formatted + " | - | - | requestUriReceived | POST - /v1/r-customer-code-information")
	/*LOG FIM*/

	customer, err := p.apps.CodeInformation.GetCodeInformation(ctx, requestCodeInformation)
	if err != nil {
		log.Println("api.codeinformation.codeinformation.codeinformation.get_code_information", err.Error())
		return c.Status(http.StatusBadRequest).JSON(codeinformationModel.CodeInformationResponseError{
			Description: err.Error(),
			Provider: codeinformationModel.Provider{
				ServiceName:  "r-customer-code-information",
				ErrorCode:    "500",
				ErrorMessage: err.Error(),
			},
		})
	}

	responseCustomer, ok := customer.(*codeinformationModel.CodeInformationServiceResponse)
	if ok {
		responseAdapter, _ := json.Marshal(responseCustomer)
		log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | orch-r-customer-code-information | 1 | - | - | " + formatted + " | - | - | responseAdapter | " + string(responseAdapter))
		return c.JSON(codeinformationModel.CodeInformationResponse{
			CustomerCode: responseCustomer.Transform(),
		})
	}

	responseCustomerErr, _ := customer.(*codeinformationModel.CodeInformationResponseProviderError)
	responseAdapter, _ := json.Marshal(responseCustomerErr)

	log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | orch-r-customer-code-information | 1 | - | - | " + formatted + " | - | - | responseAdapter | " + string(responseAdapter))

	return c.Status(http.StatusBadRequest).JSON(codeinformationModel.CodeInformationResponseError{
		Description: responseCustomerErr.Body.Fault.Faultstring.Text,
		Provider: codeinformationModel.Provider{
			ServiceName:  "r-customer-code-information",
			ErrorCode:    responseCustomerErr.Body.Fault.Faultcode.Ns0,
			ErrorMessage: responseCustomerErr.Body.Fault.Faultstring.Text,
		},
	})
}
