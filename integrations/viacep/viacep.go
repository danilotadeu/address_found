package viacep

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
	Connect(ctx context.Context, zip string) (*address.ResponseViaCep, error)
}

type viaCep struct {
	viaCepUrl string
}

func NewViaCep(viaCepUrl string) Integrations {
	return &viaCep{
		viaCepUrl: viaCepUrl,
	}
}

func (v *viaCep) Connect(ctx context.Context, zip string) (*address.ResponseViaCep, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/%s", v.viaCepUrl, zip, "json"), nil)
	if err != nil {
		log.Println("integrations.viacep.Connect.newRequest", err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("integrations.viacep.Connect.findAddress.Do", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("integrations.viacep.Connect.readAll", err.Error())
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var responseViaCep address.ResponseViaCep
		err := json.Unmarshal(respBody, &responseViaCep)
		if err != nil {
			log.Println("integrations.viacep.Connect.jsonUnmarshal", err.Error())
			return nil, err
		}
		return &responseViaCep, nil
	}
	return nil, errors.New("invalid request")
}
