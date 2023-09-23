package turno

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type turnoRepository struct {
	db *sql.DB
}

func NewTurnoRepository(database *sql.DB) TurnoRepository {
	return &turnoRepository{
		db: database,
	}
}

// FindAll implements TurnoRepository.
func (r *turnoRepository) FindAll(ctx context.Context) ([]Turno, error) {
	query := `SELECT id, nombre_paciente, nombre_matricula, fecha_turno
	FROM go_server.turno;`
	statement, err := r.db.Query(query)
	if err != nil {
		return []Turno{}, errors.New("error preparing statement")
	}
	defer statement.Close()
	var turnos []Turno
	for statement.Next() {
		var fechaTurnoStr string
		var turno Turno
		err := statement.Scan(&turno.Id, &turno.NombrePaciente, &turno.NombreMatricula, &fechaTurnoStr)
		if err != nil {
			return nil, err
		}
		fechaTurno, err := time.Parse("2006-01-02 15:04:05", fechaTurnoStr)
		if err != nil {
			return nil, err
		}
		turno.FechaTurno = &fechaTurno

		turnos = append(turnos, turno)
	}
	if err := statement.Err(); err != nil {
		return nil, err
	}
	return turnos, nil
}

// Save implements TurnoRepository.
func (r *turnoRepository) Save(ctx context.Context, turno Turno, fecha Fecha) (Turno, error) {

	query := `INSERT INTO go_server.turno(nombre_paciente,nombre_matricula,fecha_turno)
	VALUES(?,?,?)
	`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return Turno{}, errors.New("error preparing statement")
	}
	defer statement.Close()

	fechaTurno := time.Date(fecha.Year, fecha.Mes, fecha.Day, fecha.Hour, fecha.Minute, 0, 0, time.UTC)

	turno.FechaTurno = &fechaTurno

	result, err := statement.Exec(
		turno.NombrePaciente,
		turno.NombreMatricula,
		turno.FechaTurno,
	)

	if err != nil {
		return Turno{}, errors.New("error executing statement")
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return Turno{}, errors.New("error insert lastId")
	}

	turno.Id = lastId

	return turno, nil
}
