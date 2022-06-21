package address

// var ErrorAccountNotFound = errors.New("Account not found")
// var ErrorAccountExists = errors.New("Account yet exists")
// var ErrorAccountListIsEmpty = errors.New("Account list is empty")

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

//ZipResponse is a struct to response zip
type ZipResponse struct {
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

func TransformResult(response *ResponseViaCep) ZipResponse {
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
