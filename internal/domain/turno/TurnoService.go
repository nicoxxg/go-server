package turno

import (
	"context"
	"errors"
	"log"
)

type turnoService struct {
	turnoRepository TurnoRepository
}

type TurnoService interface {
	FindAll(ctx context.Context) ([]Turno, error)
	Save(ctx context.Context, requestTurno RequestTurno) (Turno, error)
}

func NewTurnoService(turnoRepository TurnoRepository) TurnoService {
	return &turnoService{
		turnoRepository: turnoRepository,
	}
}

// FindAll implements TurnoService.
func (s *turnoService) FindAll(ctx context.Context) ([]Turno, error) {
	turnos, err := s.turnoRepository.FindAll(ctx)
	if err != nil {
		log.Println("error al traer turnos desde el service", err.Error())
		return []Turno{}, err
	}
	return turnos, nil
}

// Save implements TurnoService.
func (s *turnoService) Save(ctx context.Context, requestTurno RequestTurno) (Turno, error) {

	turno := requestToTurno(requestTurno)

	fecha := requestToFecha(requestTurno)

	response, err := s.turnoRepository.Save(ctx, turno, fecha)

	if err != nil {
		return Turno{}, errors.New("error en servicio: metodo save")
	}
	return response, nil

}

func requestToTurno(requestTurno RequestTurno) Turno {
	var turno Turno
	turno.NombreMatricula = requestTurno.NombreMatricula
	turno.NombrePaciente = requestTurno.NombrePaciente
	return turno
}
func requestToFecha(requestTurno RequestTurno) Fecha {
	var fecha Fecha
	fecha.Year = requestTurno.Fecha.Year
	fecha.Day = requestTurno.Fecha.Day
	fecha.Mes = requestTurno.Fecha.Mes
	fecha.Hour = requestTurno.Fecha.Hour
	fecha.Minute = requestTurno.Fecha.Minute
	return fecha
}
