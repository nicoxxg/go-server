package turno

import "time"

type Fecha struct {
	Year   int        `json:"year"`
	Mes    time.Month `json:"mes"`
	Day    int        `json:"day"`
	Hour   int        `json:"hour"`
	Minute int        `json:"minute"`
}
