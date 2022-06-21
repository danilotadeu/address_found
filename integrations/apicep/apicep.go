package apicep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/danilotadeu/address_found/model/address"
)

type Integrations interface {
	Connect(ctx context.Context, zip string) (*address.ResponseApiCep, error)
}

type apiCep struct {
	apiCepUrl string
}

func NewApiCep(apiCepUrl string) Integrations {
	return &apiCep{
		apiCepUrl: apiCepUrl,
	}
}

func (v *apiCep) Connect(ctx context.Context, zip string) (*address.ResponseApiCep, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", v.apiCepUrl, "code=", zip), nil)
	if err != nil {
		log.Println("integrations.apicep.Connect.newRequest", err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("integrations.apicep.Connect.Do", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("integrations.apicep.Connect.readAll", err.Error())
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var responseApiCep address.ResponseApiCep
		err := json.Unmarshal(respBody, &responseApiCep)
		if err != nil {
			log.Println("integrations.apicep.Connect.jsonUnmarshal", err.Error())
			return nil, err
		}
		return &responseApiCep, nil
	}
	return nil, errors.New("invalid request")
}
