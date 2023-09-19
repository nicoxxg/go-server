package cliente

import (
	"context"
	"database/sql"

	"github.com/nicoxxg/go-server/internal/domain/cliente"
)

type clienteRepository struct {
	db *sql.DB
}

func NewClienteRepository(database *sql.DB) ClienteRepository {
	return &clienteRepository{
		db: database,
	}
}

// findAll implements ClienteRepository.
func (*clienteRepository) findAll(ctx context.Context) ([]cliente.Cliente, error) {
	panic("unimplemented")
}

// findByEmail implements ClienteRepository.
func (*clienteRepository) findByEmail(ctx context.Context, email string) (cliente.Cliente, error) {
	panic("unimplemented")
}

// findById implements ClienteRepository.
func (*clienteRepository) findById(ctx context.Context, id string) (cliente.Cliente, error) {
	panic("unimplemented")
}

// save implements ClienteRepository.
func (*clienteRepository) save(ctx context.Context, cliente cliente.Cliente) error {
	panic("unimplemented")
}

// update implements ClienteRepository.
func (*clienteRepository) update(ctx context.Context, cliente cliente.Cliente) error {
	panic("unimplemented")
}
