package cliente

import (
	"context"
	"log"
)

type clienteService struct {
	clienteRepository ClienteRepository
}

type ClientService interface {
	FindAll(ctx context.Context) ([]Cliente, error)
	FindById(ctx context.Context, id string) (Cliente, error)
	FindByEmail(ctx context.Context, email string) (Cliente, error)
	Save(ctx context.Context, cliente Cliente) error
	Update(ctx context.Context, cliente Cliente) error
}

func NewClientService(clienteRepository ClienteRepository) ClientService {
	return &clienteService{
		clienteRepository: clienteRepository,
	}
}

// findAll implements ClientService.
func (s *clienteService) FindAll(ctx context.Context) ([]Cliente, error) {
	productos, err := s.clienteRepository.findAll(ctx)
	if err != nil {
		log.Println("error al traer clientes desde el service", err.Error())
		return []Cliente{}, ErrEmptyList
	}
	return productos, nil
}

// findByEmail implements ClientService.
func (s *clienteService) FindByEmail(ctx context.Context, email string) (Cliente, error) {
	panic("unimplemented")
}

// findById implements ClientService.
func (s *clienteService) FindById(ctx context.Context, id string) (Cliente, error) {
	panic("unimplemented")
}

// save implements ClientService.
func (s *clienteService) Save(ctx context.Context, cliente Cliente) error {
	panic("unimplemented")
}

// update implements ClientService.
func (s *clienteService) Update(ctx context.Context, cliente Cliente) error {
	panic("unimplemented")
}
