package cliente

import (
	"context"
	"errors"
)

type ClienteRepository interface {
	FindAll(ctx context.Context) ([]Cliente, error)
	FindById(ctx context.Context, id int64) (Cliente, error)
	FindByEmail(ctx context.Context, email string) (Cliente, error)
	Save(ctx context.Context, cliente Cliente) (Cliente, error)
	Update(ctx context.Context, cliente Cliente) (Cliente, error)
}

var (
	ErrEmptyList = errors.New("la lista de productos esta vacia")
	ErrNotFound  = errors.New("producto no encontrado")
)
