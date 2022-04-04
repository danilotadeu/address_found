package codeinformation

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	codeinformationModel "github.com/danilotadeu/r-customer-code-information/model/codeinformation"
)

const (
	contentType      = "Content-Type"
	contentTypeValue = "text/xml"
)

//App is a contract to CodeInformation..
type App interface {
	GetCodeInformation(ctx context.Context, requestCodeInformation *codeinformationModel.CodeInformationRequest) (interface{}, error)
}

type appImpl struct {
	urlProvider  string
	portProvider string
}

//NewApp init a codeInformation
func NewApp(urlProvider, portProvider string) App {
	return &appImpl{
		urlProvider:  urlProvider,
		portProvider: portProvider,
	}
}

//GetCodeInformation get a information code in webservice...
func (a appImpl) GetCodeInformation(ctx context.Context, requestCodeInformation *codeinformationModel.CodeInformationRequest) (interface{}, error) {
	requestbody := prepareBody(requestCodeInformation)
	body, err := xml.Marshal(requestbody)
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.xml_marshal", err.Error())
		return nil, err
	}

	url := fmt.Sprintf("http://%s:%s/v1/r-customer-code-information-service", a.urlProvider, a.portProvider)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(string(body)))

	req.Header.Add("Content-Type", "text/xml")
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.new_request", err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	client := http.DefaultClient

	/* LOG INICIO */
	now := time.Now()
	formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	log.Println("    INFO    EngDB | Engineering | GMID |  | - | ENGNB003153 | rCustomerCodeInformationService | 1 | - | - | " + formatted + " | - | - | [rCustomerCodeInformationService] |  input: {'msisdn':'123'}")
	timeDur := 5 * time.Second
	log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | rCustomerCodeInformationService | 1 | - | - | " + formatted + " | - | - | Http Provider request | TIMEOUT[" + strconv.Itoa(int(timeDur.Seconds())) + "] METHOD[POST] URL[" + url + "]")
	payload, err := json.Marshal(requestbody)
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.json_marshal_payload", err.Error())
		return nil, err
	}
	log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | rCustomerCodeInformationService | 1 | - | - | " + formatted + " | - | - | request body | " + string(payload))
	log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | rCustomerCodeInformationService | 1 | - | - | " + formatted + " | - | - | provider header | [" + contentType + ", " + contentTypeValue + "]")
	/* LOG FIM */
	resp, err := client.Do(req)
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.do", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.ioutil_readall", err.Error())
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var responseService codeinformationModel.CodeInformationServiceResponse
		err = xml.Unmarshal(data, &responseService)
		responseProviderJSON, _ := json.Marshal(responseService)
		log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | rCustomerCodeInformationService | 1 | - | - | " + formatted + " | - | - | out | " + string(responseProviderJSON))
		if err != nil {
			log.Println("app.codeinformation.codeinformation.codeinformation.xml_unmarshal", err.Error())
			return nil, err
		}

		return &responseService, nil
	}

	var responseService codeinformationModel.CodeInformationResponseProviderError
	err = xml.Unmarshal(data, &responseService)
	responseProviderJSON, err := json.Marshal(responseService)
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.json_marshal_response_provider_json", err.Error())
		return nil, err
	}
	log.Println("    DEBUG    EngDB | Engineering | GMID |  | - | ENGNB003153 | rCustomerCodeInformationService | 1 | - | - | " + formatted + " | - | - | out | " + string(responseProviderJSON))
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.xml_unmarshal_provider_error", err.Error())
		return nil, err
	}
	return &responseService, nil
}

func prepareBody(data *codeinformationModel.CodeInformationRequest) codeinformationModel.CustomerRetrieveRequest {
	requestbody := codeinformationModel.CustomerRetrieveRequest{}
	requestbody.Header.Security.UsernameToken.Username = "ADMX"
	requestbody.Header.Security.UsernameToken.Password.Type = data.Password
	if val, err := strconv.Atoi(data.Customer.ID); err == nil {
		requestbody.Body.CustomerRetrieveRequest.InputAttributes.CustomerRead.CsId = val
		requestbody.Body.CustomerRetrieveRequest.InputAttributes.PaymentArrangementsRead.CsId = val
		requestbody.Body.CustomerRetrieveRequest.InputAttributes.AddressesRead.CsId = val
		requestbody.Body.CustomerRetrieveRequest.InputAttributes.CustomerInfoRead.CsId = val
	}
	requestbody.Body.CustomerRetrieveRequest.InputAttributes.CustomerRead.SyncWithDb = data.Customer.SyncFlag
	requestbody.Body.CustomerRetrieveRequest.SessionChangeRequest.Values.Item.Key = data.UserID
	return requestbody
}
