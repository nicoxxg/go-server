package cliente

import (
	"context"
	"errors"
)

type ClienteRepository interface {
	findAll(ctx context.Context) ([]Cliente, error)
	findById(ctx context.Context, id int64) (Cliente, error)
	findByEmail(ctx context.Context, email string) (Cliente, error)
	save(ctx context.Context, cliente Cliente) error
	update(ctx context.Context, cliente Cliente) error
}

var (
	ErrEmptyList = errors.New("la lista de productos esta vacia")
	ErrNotFound  = errors.New("producto no encontrado")
)
