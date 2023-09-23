package turno

import "context"

type TurnoRepository interface {
	FindAll(ctx context.Context) ([]Turno, error)
	Save(ctx context.Context, turno Turno, fecha Fecha) (Turno, error)
}
