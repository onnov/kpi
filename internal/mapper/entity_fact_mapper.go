package mapper

import "net/http"

// EntityFactMapper проверяет запрос и формирует json в соответствии с типами полей
func EntityFactMapper(r *http.Request) (*[]byte, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	// Описание формы, поля и их типы
	fields := map[string]string{
		"period_start":            "date",
		"period_end":              "date",
		"period_key":              "str",
		"indicator_to_mo_id":      "int",
		"indicator_to_mo_fact_id": "int",
		"value":                   "int",
		"fact_time":               "date",
		"is_plan":                 "int",
		"auth_user_id":            "int",
		"comment":                 "str",
	}

	return formToJson(r, fields)
}
