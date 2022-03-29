package codeinformation

import (
	"context"
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

//App is a contract to CodeInformation..
type App interface {
	GetCodeInformation(ctx context.Context, requestCodeInformation *codeinformationModel.CodeInformationRequest) (*codeinformationModel.CodeInformationServiceResponse, error)
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
func (a appImpl) GetCodeInformation(ctx context.Context, requestCodeInformation *codeinformationModel.CodeInformationRequest) (*codeinformationModel.CodeInformationServiceResponse, error) {
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

	//LOG AQUI
	url := fmt.Sprintf("http://%s:%s/v1/r-customer-code-information-service", a.urlProvider, a.portProvider)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(body)))
	req.Header.Add("Content-Type", "text/xml")
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.new_request", err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	client := http.DefaultClient
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

	/*
		fazer unmarshal com map interface e verificar se existe tag
	*/

	var responseService codeinformationModel.CodeInformationServiceResponse
	err = xml.Unmarshal(data, &responseService)
	if err != nil {
		log.Println("app.codeinformation.codeinformation.codeinformation.xml_unmarshal", err.Error())
		return nil, err
	}

	return &responseService, nil
}
