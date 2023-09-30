package cliente

type Cliente struct {
	Id         int64  `json:"id"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Email      string `json:"email"`
	Contrasena string `json:"contrasena"`
	Activo     bool   `json:"activo"`
}
type ClientJson struct {
	Email     string
	Password  string
	Authority string
	Exp       int64
}

type ClientRequest struct {
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Email      string `json:"email"`
	Contrasena string `json:"contraseña"`
}
