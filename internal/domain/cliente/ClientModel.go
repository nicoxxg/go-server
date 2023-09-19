package cliente

type Cliente struct {
	Id         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Email      string `json:"email"`
	Contrasena string `json:"contrasena"`
	Activo     bool   `json:"activo"`
}
