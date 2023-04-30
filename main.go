package main

import (
	"fmt"

	"github.com/fepc18/godesde0/variables"
)

func main() {
	var x int
	x = 5
	fmt.Println(x)

	fmt.Println("**Metodos de variables Enteras**")
	variables.MuestroEnteros()

	fmt.Println("**Metodos de variables Resto**")
	variables.RestoVariables()

	fmt.Println("**Convertir a texto**")
	estado, texto := variables.ConviertoaTexto(10)
	fmt.Println(estado, texto)

}
