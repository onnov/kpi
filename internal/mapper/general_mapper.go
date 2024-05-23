package mapper

import (
	"encoding/json"
	"errors"
	"kpi/internal/helper"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// formToJson проверяет форму и формирует json
func formToJson(r *http.Request, fields map[string]string) (*[]byte, error) {
	p := make(map[string]interface{})

	for k, t := range fields {
		switch t {
		case "date":
			re := regexp.MustCompile("^\\d{4}\\-\\d{2}\\-\\d{2}$")
			v := strings.TrimSpace(r.FormValue(k))
			if re.MatchString(v) {
				date, err := time.Parse("2006-01-02", v)
				if err != nil {
					return nil, err
				}
				p[k] = date
			} else {
				return nil, errors.New("The '" + k + "' field must be in the format YYYY-MM-DD")
			}
			break
		case "str":
			p[k] = r.FormValue(k)
			break
		case "int":
			re := regexp.MustCompile("[0-9]")
			v := strings.TrimSpace(r.FormValue(k))
			if re.MatchString(v) {
				p[k] = helper.Sti(v)
			} else {
				return nil, errors.New("Field '" + k + "' must be an integer")
			}
			break
		}
	}
	jsonString, err := json.Marshal(&p)
	if err != nil {
		return nil, err
	}

	return &jsonString, nil
}
