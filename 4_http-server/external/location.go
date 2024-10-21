package external

import (
	"encoding/json"
	"io"
	"net/http"
)

type UFData struct {
	Id      int    `json:"id"`
	Name    string `json:"nome"`
	Acronym string `json:"sigla"`
}

func GetUfsLocationApi() []UFData {
	response, err := http.Get("https://servicodados.ibge.gov.br/api/v1/localidades/estados?orderBy=nome")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bodyAsBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var ufsList []UFData

	err = json.Unmarshal(bodyAsBytes, &ufsList)

	if err != nil {
		panic(err)
	}

	return ufsList
}
