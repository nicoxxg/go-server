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
	FindById(ctx context.Context, id int64) (Cliente, error)
	FindByEmail(ctx context.Context, email string) (Cliente, error)
	Save(ctx context.Context, cliente Cliente) error
	Update(ctx context.Context, cliente Cliente) error
}

func NewClientService(clienteRepository ClienteRepository) ClientService {
	return &clienteService{
		clienteRepository: clienteRepository,
	}
}

// findByEmail implements ClientService.
func (s *clienteService) FindByEmail(ctx context.Context, email string) (Cliente, error) {

	cliente, err := s.clienteRepository.findByEmail(ctx, email)

	if err != nil {
		log.Println("error al traer cliente desde el service", err.Error())
		return Cliente{}, err
	}

	return cliente, nil

}

// findAll implements ClientService.

func (s *clienteService) FindById(ctx context.Context, id int64) (Cliente, error) {
	cliente, err := s.clienteRepository.findById(ctx, id)

	if err != nil {
		log.Println("error al traer cliente desde el service", err.Error())
		return Cliente{}, err
	}
	return cliente, nil
}
func (s *clienteService) FindAll(ctx context.Context) ([]Cliente, error) {
	productos, err := s.clienteRepository.findAll(ctx)
	if err != nil {
		log.Println("error al traer clientes desde el service", err.Error())
		return []Cliente{}, ErrEmptyList
	}
	return productos, nil
}

// findById implements ClientService.

// save implements ClientService.
func (s *clienteService) Save(ctx context.Context, cliente Cliente) error {
	panic("unimplemented")
}

// update implements ClientService.
func (s *clienteService) Update(ctx context.Context, cliente Cliente) error {
	panic("unimplemented")
}
