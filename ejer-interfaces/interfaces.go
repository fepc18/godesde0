package ejerinterfaces

import (
	"fmt"

	"github.com/fepc18/godesde0/interfaces"
)

func HumanoRespirando(h interfaces.Humano) {
	h.Respirar()
	fmt.Printf("Soy un/a %s esta respirando \n", h.Sexo())
}
