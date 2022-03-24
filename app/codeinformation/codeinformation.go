package codeinformation

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	codeinformationModel "github.com/danilotadeu/r-customer-code-information/model/codeinformation"
	"github.com/valyala/fasthttp"
)

//App is a contract to CodeInformation..
type App interface {
	GetCodeInformation(ctx *fasthttp.RequestCtx, requestCodeInformation *codeinformationModel.CodeInformationRequest, clientID, messageID string) (*string, error)
}

type appImpl struct{}

//NewApp init a codeInformation
func NewApp() App {
	return &appImpl{}
}

//GetCodeInformation get a information code in webservice...
func (a appImpl) GetCodeInformation(ctx *fasthttp.RequestCtx, requestCodeInformation *codeinformationModel.CodeInformationRequest, clientID, messageID string) (*string, error) {
	/*
		TODO: Fazer o XML de envio e tratar erros de response
		TODO: Mudar para requisição POST e envia body
	*/
	resp, err := http.Get("http://127.0.0.1:4000/v1/r-customer-code-information-service")
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

	return &responseService.CustomerRead.CsCode, nil
}
