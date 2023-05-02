package modelos

type Mujer struct {
	Hombre //Herencia de Hombre
}

// Como Mujer hereda de Hombre, no es necesario declarar los métodos de Hombre
// Sobreescritura de métodos
func (m *Mujer) Sexo() string { return "Mujer" }
