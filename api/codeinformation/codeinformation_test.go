package codeinformation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/danilotadeu/r-customer-code-information/app"
	"github.com/danilotadeu/r-customer-code-information/model/codeinformation"
	"github.com/danilotadeu/r-customer-code-information/tests_personal"

	"github.com/gofiber/fiber/v2"
)

func Test_apiImpl_CodeInformationHandler(t *testing.T) {
	os.Setenv("URL_PROVIDER", "localhost")
	os.Setenv("PORT_PROVIDER", "4000")

	apps := app.Register()
	api := apiImpl{
		apps: apps,
	}
	tests := []struct {
		name        string
		method      string
		contentType string
		header      map[string]string
		url         string
		urlReq      string
		handlerFunc func(c *fiber.Ctx) error
		Input       codeinformation.CodeInformationRequest
		want        int
		bodyShow    bool
	}{
		{
			name:        "r-customer-code-information",
			method:      "POST",
			contentType: "application/json",
			header:      map[string]string{"messageId": "123456", "Clientid": "123"},
			url:         "/v1/r-customer-code-information",
			urlReq:      "/v1/r-customer-code-information",
			handlerFunc: api.CodeInformationHandler,
			Input: codeinformation.CodeInformationRequest{
				UserID:   "F8036589",
				Password: "ADMX",
				Customer: struct {
					ID       string "json:\"id\""
					SyncFlag bool   "json:\"syncFlag\""
				}{
					ID:       "116084470",
					SyncFlag: false,
				},
			},
			want:     200,
			bodyShow: true,
		},
		{
			name:        "r-customer-code-information",
			method:      "POST",
			contentType: "application/json",
			header:      map[string]string{"messageId": "123456", "Clientid": "123"},
			url:         "/v1/r-customer-code-information",
			urlReq:      "/v1/r-customer-code-information",
			handlerFunc: api.CodeInformationHandler,
			Input: codeinformation.CodeInformationRequest{
				UserID: "F8036589",
				// Password: "ADMX",
				Customer: struct {
					ID       string "json:\"id\""
					SyncFlag bool   "json:\"syncFlag\""
				}{
					ID:       "116084470",
					SyncFlag: false,
				},
			},
			want:     400,
			bodyShow: true,
		},
	}

	for _, tt := range tests {
		requestBody, err := json.Marshal(&tt.Input)
		if err != nil {
			t.Errorf("Error json.Marshal:%s", err.Error())
			return
		}
		t.Run(tt.name, func(t *testing.T) {
			if tt.bodyShow {
				fmt.Println("Json: ", string(requestBody))
			}
			tests_personal.TestNewRequest(t, tt.url, tt.urlReq, tt.method,
				tt.handlerFunc,
				bytes.NewBuffer(requestBody),
				tt.contentType, tt.header, tt.want, tt.bodyShow)
		})
	}
}
