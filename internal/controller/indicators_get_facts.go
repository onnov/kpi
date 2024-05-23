package controller

import (
	"encoding/json"
	"kpi/internal/mapper"
	"kpi/internal/repository"
	"net/http"
)

// IndicatorsGetFactsHandler обработчик запроса /_api/indicators/get_facts
func IndicatorsGetFactsHandler(w http.ResponseWriter, r *http.Request) {
	dto, err := mapper.DtoFactReqMapper(r)
	if err != nil {
		setError(w, err)
		return
	}

	db := repository.GetDb()
	res, err := db.GetFacts(dto)
	if err != nil {
		setError(w, err)
		return
	}

	jsonRes, err := json.Marshal(&res)
	if err != nil {
		setError(w, err)
		return
	}

	set200ok(w, string(jsonRes))
}
