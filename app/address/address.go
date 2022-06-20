package address

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/danilotadeu/address_found/model/address"
)

//App is a contract to Address..
type App interface {
	FindAddress(ctx context.Context, zip string) (*address.ResponseViaCep, error)
}

type appImpl struct {
	urlViaCep string
}

//NewApp init a Address
func NewApp(urlViaCep string) App {
	return &appImpl{
		urlViaCep: urlViaCep,
	}
}

//FindAddress find a address in viacep or in another integrations..
func (a *appImpl) FindAddress(ctx context.Context, zip string) (*address.ResponseViaCep, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/%s", a.urlViaCep, zip, "json"), nil)
	if err != nil {
		log.Println("app.address.findAddress.newRequest", err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("app.address.findAddress.Do", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("app.address.findAddress.readAll", err.Error())
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var responseViaCep address.ResponseViaCep
		err := json.Unmarshal(respBody, &responseViaCep)
		if err != nil {
			log.Println("app.address.findAddress.jsonUnmarshal", err.Error())
			return nil, err
		}
		return &responseViaCep, nil
	}
	return &address.ResponseViaCep{}, nil
}
