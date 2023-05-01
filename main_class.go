package main_class

import (
	"fmt"
	"runtime"

	"github.com/fepc18/godesde0/variables"
)

func main_class() {
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

	//Condicionales
	os := runtime.GOOS
	fmt.Println("El sistema operativo es: ", os)
	if os == "windows" {
		fmt.Println("Es windows")
	} else if os == "darwin" {
		fmt.Println("Es Mac")
	} else {
		fmt.Println("Es otro")
	}

	if os2 := runtime.GOOS; os2 == "windows" || os2 == "darwin" {
		fmt.Println("Es windows o mac")
	} else {
		fmt.Println("Es otro")
	}

	//Switch
	switch os3 := runtime.GOOS; os3 {
	case "windows":
		fmt.Println("Es windows")
	case "darwin":
		fmt.Println("Es Mac")
	default:
		fmt.Printf("%s \n", os3)
	}

}
