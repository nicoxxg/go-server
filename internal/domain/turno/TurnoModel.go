package turno

import "time"

type Turno struct {
	Id              int64      `json:"id"`
	NombrePaciente  string     `json:"nombre"`
	NombreMatricula string     `json:"nombreMatricula"`
	FechaTurno      *time.Time `json:"fechaTurno"`
}

type RequestTurno struct {
	NombrePaciente  string `json:"nombrePaciente"`
	NombreMatricula string `json:"nombreMatricula"`
	Fecha           Fecha  `json:"fecha"`
}
