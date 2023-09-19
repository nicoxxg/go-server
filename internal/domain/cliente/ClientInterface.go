package cliente

import (
	"context"
)

type ClienteRepository interface {
	findAll(ctx context.Context) ([]Cliente, error)
	findById(ctx context.Context, id string) (Cliente, error)
	findByEmail(ctx context.Context, email string) (Cliente, error)
	save(ctx context.Context, cliente Cliente) error
	update(ctx context.Context, cliente Cliente) error
}
