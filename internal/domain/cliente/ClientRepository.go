package cliente

import (
	"context"
	"database/sql"
	"errors"
)

type clienteRepository struct {
	db *sql.DB
}

func NewClienteRepository(database *sql.DB) ClienteRepository {
	return &clienteRepository{
		db: database,
	}
}

// update implements ClienteRepository.
func (r *clienteRepository) update(ctx context.Context, cliente Cliente) (Cliente, error) {
	query := `
        UPDATE go_server.cliente
        SET nombre = ?, apellido = ?, email = ?, contraseña = ?, activo = ?
        WHERE id = ?
    `
	statement, err := r.db.Prepare(query)

	if err != nil {
		return Cliente{}, errors.New("error preparing statement")
	}

	defer statement.Close()

	result, err := statement.Exec(
		cliente.Nombre,
		cliente.Apellido,
		cliente.Email,
		cliente.Contrasena,
		cliente.Activo,
		cliente.Id, // Suponiendo que tengas el campo "Id" en tu struct Cliente
	)

	if err != nil {
		return Cliente{}, errors.New("error executing statement")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Cliente{}, errors.New("error getting rows affected")
	}

	if rowsAffected == 0 {
		return Cliente{}, errors.New("no rows updated")
	}

	return cliente, nil
}

// save implements ClienteRepository.

func (r *clienteRepository) save(ctx context.Context, cliente Cliente) (Cliente, error) {

	query := `INSERT INTO go_server.cliente(nombre, apellido, email,contraseña, activo)
	VALUES(?,?,?,?,?)
	`
	statement, err := r.db.Prepare(query)
	if err != nil {
		return Cliente{}, errors.New("error preparing statement")
	}
	defer statement.Close()
	cliente.Activo = true

	result, err := statement.Exec(
		cliente.Nombre,
		cliente.Apellido,
		cliente.Email,
		cliente.Contrasena,
		cliente.Activo,
	)

	if err != nil {
		return Cliente{}, errors.New("error executing statement")
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return Cliente{}, errors.New("error insert lastId")
	}

	cliente.Id = lastId

	return cliente, nil
}

// findByEmail implements ClienteRepository.

func (r *clienteRepository) findByEmail(ctx context.Context, email string) (Cliente, error) {

	query := `SELECT id, nombre, apellido, email, activo
	FROM go_server.cliente
	WHERE email = ?;
	`

	var cliente Cliente

	err := r.db.QueryRow(query, email).Scan(&cliente.Id, &cliente.Nombre, &cliente.Apellido, &cliente.Email, &cliente.Activo)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Cliente{}, errors.New("Cliente No Encontrado")
		}

		return Cliente{}, nil
	}

	return cliente, nil

}

// findById implements ClienteRepository.
func (r *clienteRepository) findById(ctx context.Context, id int64) (Cliente, error) {
	query := `SELECT id, nombre, apellido, email, activo
	FROM go_server.cliente
	WHERE id = ?;
	`

	var cliente Cliente
	err := r.db.QueryRow(query, id).Scan(&cliente.Id, &cliente.Nombre, &cliente.Apellido, &cliente.Email, &cliente.Activo)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Cliente{}, errors.New("Cliente No Encontrado")
		}

		return Cliente{}, nil
	}

	return cliente, nil
}

// findAll implements ClienteRepository.
func (r *clienteRepository) findAll(ctx context.Context) ([]Cliente, error) {
	query := `SELECT id, nombre, apellido, email, activo
	FROM go_server.cliente;`

	statement, err := r.db.Query(query)
	if err != nil {
		return []Cliente{}, errors.New("error preparing statement")
	}
	defer statement.Close()

	var clientes []Cliente

	for statement.Next() {
		var cliente Cliente
		err := statement.Scan(&cliente.Id, &cliente.Nombre, &cliente.Apellido, &cliente.Email, &cliente.Activo)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}

	if err := statement.Err(); err != nil {
		return nil, err
	}

	return clientes, nil

}
