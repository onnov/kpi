package dto

import "time"

type GetFactsRequest struct {
	PeriodStart     time.Time `json:"period_start"`
	PeriodEnd       time.Time `json:"period_end"`
	PeriodKey       string    `json:"period_key"`
	IndicatorToMoId int       `json:"indicator_to_mo_id"`
}
