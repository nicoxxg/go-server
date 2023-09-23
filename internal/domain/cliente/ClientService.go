package cliente

import (
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type clienteService struct {
	clienteRepository ClienteRepository
}

type ClientService interface {
	FindAll(ctx context.Context) ([]Cliente, error)
	FindById(ctx context.Context, id int64) (Cliente, error)
	FindByEmail(ctx context.Context, email string) (Cliente, error)
	SaveClient(ctx context.Context, cliente ClientRequest) (Cliente, error)
	Update(ctx context.Context, cliente Cliente) (Cliente, error)
}

func NewClientService(clienteRepository ClienteRepository) ClientService {
	return &clienteService{
		clienteRepository: clienteRepository,
	}
}

func (s *clienteService) SaveClient(ctx context.Context, cliente ClientRequest) (Cliente, error) {
	contraseñaEncriptada, err := bcrypt.GenerateFromPassword([]byte(cliente.Contrasena), bcrypt.DefaultCost)

	if err != nil {
		return Cliente{}, errors.New("eror al hashear la contraseña")
	}
	cliente.Contrasena = string(contraseñaEncriptada)

	clienteTransforado := requestToClient(cliente)

	response, err := s.clienteRepository.save(ctx, clienteTransforado)

	if err != nil {
		return Cliente{}, errors.New("error en servicio: metodo Post")
	}

	return response, nil
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

// update implements ClientService.
func (s *clienteService) Update(ctx context.Context, cliente Cliente) (Cliente, error) {
	panic("unimplemented")
}
func requestToClient(clientRequest ClientRequest) Cliente {
	var cliente Cliente

	cliente.Nombre = clientRequest.Nombre
	cliente.Apellido = clientRequest.Apellido
	cliente.Email = clientRequest.Email
	cliente.Contrasena = clientRequest.Contrasena

	return cliente
}
