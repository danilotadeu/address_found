package tests_personal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T, url, urlReq string,
	method string,
	handlerFunc func(c *fiber.Ctx) error,
	data *bytes.Buffer, contentType string,
	header map[string]string,
	want int, bodyShow bool) {

	app := fiber.New()

	if method == "POST" {
		app.Post(url, handlerFunc)
	} else if method == "PUT" {
		app.Put(url, handlerFunc)
	} else if method == "GET" {
		app.Get(url, handlerFunc)
	} else if method == "DELETE" {
		app.Delete(url, handlerFunc)
	} else if method == "PATCH" {
		app.Patch(url, handlerFunc)
	}

	var req *http.Request
	if data == nil {
		req = httptest.NewRequest(method, urlReq, nil)
	} else {
		req = httptest.NewRequest(method, urlReq, data)
	}

	if len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := app.Test(req)
	if err != nil {
		if want == 400 {
			t.Logf("app.Test:%s", err.Error())
			return
		}
		t.Errorf("Error app.Test:%s", err.Error())
		return
	}

	utils.AssertEqual(t, nil, err, "app.Test(req)")
	assert.Equal(t, want, resp.StatusCode)

	if bodyShow {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error ioutil.ReadAll:%s", err.Error())
			return
		}

		t.Log("status:", resp.StatusCode)
		t.Log("\nResp :\n", string(body))
	}
}
