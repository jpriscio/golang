package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, url := range os.Args[1:] {
		req, err := http.Get(url)

		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		var seucep ViaCEP

		json.Unmarshal(res, &seucep)

		file, err := os.Create("CEP.txt")
		if err != nil {
			panic("Erro ao criar arquivo")
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nEstado: %s\nUF: %s\n", seucep.Cep, seucep.Estado, seucep.Uf))
		if err != nil {
			panic(err)
		}

	}

}
