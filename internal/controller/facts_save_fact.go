package controller

import (
	"kpi/internal/mapper"
	"kpi/internal/repository"
	"net/http"
)

// FactsSaveFactHandler обработчик запроса /_api/facts/save_fact
func FactsSaveFactHandler(w http.ResponseWriter, r *http.Request) {
	j, err := mapper.EntityFactMapper(r)
	if err != nil {
		setError(w, err)
		return
	}

	q := repository.GetQueue()
	q.Add(j)

	set200ok(w, `{"result": "ok"}`)
}
