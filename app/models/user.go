package models

type Usuario struct {
	ID        string `json:"id"`
	Nombre    string `json:"nombre"`
	ApPaterno string `json:"apPaterno"`
	ApMaterno string `json:"apMaterno"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Ciudad    string `json:"ciudad"`
	Estado    string `json:"estado"`
	Usuario   string `json:"usuario"`
	Password  string `json:"-"`
	Rol       string `json:"rol"`
	Image     string `json:"imagen"`
}
