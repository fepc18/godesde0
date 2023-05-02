package main

import (
	ejerinterfaces "github.com/fepc18/godesde0/ejer-interfaces" //Alias
	"github.com/fepc18/godesde0/modelos"
)

func main() {
	// numero, resultado := ejercicios.Funcion("dddd")
	// fmt.Println("Numero", numero)
	// fmt.Println("Resultado", resultado)

	//iteraciones.Iterar()

	// result := ejercicios.Multiplicar(5)
	// fmt.Println(result)

	//files.GrabarTabla()
	//files.SumaTabla()

	//funciones.Calculos()

	// funciones.LllamarClosure()
	// numbertoExponencial := 6
	// fmt.Println(numbertoExponencial, " exponencial 2=", funciones.Exponencial(numbertoExponencial))
	/*
		arreglosslices.MuestroArreglos()

		arreglosslices.MuestroSlice()

		arreglosslices.Capacidades()
	*/
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	Pedro := new(modelos.Hombre)
	ejerinterfaces.HumanoRespirando(Pedro)

	Maria := new(modelos.Mujer)
	ejerinterfaces.HumanoRespirando(Maria)

}
