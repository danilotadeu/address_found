package address

//ZipRequest is a struct to find zip
type ZipRequest struct {
	Zip string `json:"zip" validate:"required,min=8,max=8"`
}

//ResponseViaCep used to parse response viacep to struct
type ResponseViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

//ResponseApiCep used to parse response apicep to struct
type ResponseApiCep struct {
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	StatusText string `json:"statusText"`
}

//ZipResponse is a struct to response zip
type ZipResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade,omitempty"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge,omitempty"`
	Gia         string `json:"gia,omitempty"`
	Ddd         string `json:"ddd,omitempty"`
	Siafi       string `json:"siafi,omitempty"`
}

func TransformResultViaCep(response *ResponseViaCep) ZipResponse {
	return ZipResponse{
		Cep:         response.Cep,
		Logradouro:  response.Logradouro,
		Complemento: response.Complemento,
		Bairro:      response.Bairro,
		Localidade:  response.Localidade,
		Uf:          response.Uf,
		Ibge:        response.Ibge,
		Gia:         response.Gia,
		Ddd:         response.Ddd,
		Siafi:       response.Siafi,
	}
}

func TransformResultApiCep(response *ResponseApiCep) ZipResponse {
	return ZipResponse{
		Cep:        response.Code,
		Logradouro: response.Address,
		Bairro:     response.District,
		Localidade: response.City,
		Uf:         response.State,
	}
}
