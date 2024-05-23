package mapper

import (
	"encoding/json"
	"kpi/internal/dto"
	"net/http"
)

// DtoFactReqMapper проверяет запрос и мфппит форму в dto.GetFactsRequest
func DtoFactReqMapper(r *http.Request) (*dto.GetFactsRequest, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	// Описание формы, поля и их типы
	fields := map[string]string{
		"period_start":       "date",
		"period_end":         "date",
		"period_key":         "str",
		"indicator_to_mo_id": "int",
	}
	jsonString, err := formToJson(r, fields)
	if err != nil {
		return nil, err
	}

	var gfr dto.GetFactsRequest
	err = json.Unmarshal(*jsonString, &gfr)
	if err != nil {
		return nil, err
	}

	return &gfr, nil
}
