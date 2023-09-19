package cliente

type Cliente struct {
	id         int64  `json:"id"`
	nombre     string `json:"nombre"`
	apellido   string `json:"apellido"`
	email      string `json:"email"`
	contrasena string `json:"contrasena"`
}
