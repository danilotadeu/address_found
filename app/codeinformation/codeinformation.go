package codeinformation

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	codeinformationModel "github.com/danilotadeu/r-customer-code-information/model/codeinformation"
	"github.com/valyala/fasthttp"
)

//App is a contract to CodeInformation..
type App interface {
	GetCodeInformation(ctx *fasthttp.RequestCtx, requestCodeInformation *codeinformationModel.CodeInformationRequest, clientID, messageID string) (*string, error)
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
func (a appImpl) GetCodeInformation(ctx *fasthttp.RequestCtx, requestCodeInformation *codeinformationModel.CodeInformationRequest, clientID, messageID string) (*string, error) {
	requestbody := codeinformationModel.CustomerRetrieveRequest{}
	requestbody.Header.Security.UsernameToken.Username = "ADMX"
	requestbody.Header.Security.UsernameToken.Password.Type = "ADMX"
	if val, err := strconv.Atoi(requestCodeInformation.Customer.ID); err == nil {
		requestbody.Body.CustomerRetrieveRequest.InputAttributes.CustomerRead.CsId = val
		requestbody.Body.CustomerRetrieveRequest.InputAttributes.PaymentArrangementsRead.CsId = val
		requestbody.Body.CustomerRetrieveRequest.InputAttributes.AddressesRead.CsId = val
		requestbody.Body.CustomerRetrieveRequest.InputAttributes.CustomerInfoRead.CsId = val
	}
	requestbody.Body.CustomerRetrieveRequest.InputAttributes.CustomerRead.SyncWithDb = requestCodeInformation.Customer.SyncFlag
	requestbody.Body.CustomerRetrieveRequest.SessionChangeRequest.Values.Item.Key = requestCodeInformation.UserID

	body, err := xml.Marshal(requestbody)
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.xml_marshal", err.Error())
		return nil, err
	}
	resp, err := http.Post("http://"+a.urlProvider+":"+a.portProvider+"/v1/r-customer-code-information-service", "text/xml", strings.NewReader(string(body)))
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.body_parser", err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.ioutil_readall", err.Error())
		return nil, err
	}

	var responseService codeinformationModel.CodeInformationServiceResponse
	xml.Unmarshal(byteValue, &responseService)

	if responseService.Body.CustomerRetrieveResponse.CustomerRead.CsCode == "" {
		return nil, errors.New("cs code empty")
	}

	return &responseService.Body.CustomerRetrieveResponse.CustomerRead.CsCode, nil
}
