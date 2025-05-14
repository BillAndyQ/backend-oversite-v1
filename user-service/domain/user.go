package domain

type Role string

const (
	RoleAdmin     Role = "administrador"
	RoleContador  Role = "contador"
	RoleInspector Role = "inspector"
	RoleGerente   Role = "gerente"
)

type User struct {
	ID         uint
	Username   string
	Names      string
	Password   string
	Dni        string
	Address    string
	Phone      string
	TipoSangre string
	Role       Role `json:"role"`
}

func IsValidRole(role Role) bool {
	switch role {
	case RoleAdmin, RoleContador, RoleInspector, RoleGerente:
		return true
	default:
		return false
	}
}
