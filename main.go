package main

import (
	"fmt"

	"github.com/fepc18/godesde0/goroutines"
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

	//******************Ejercicio Interfaces******************
	/*Pedro := new(modelos.Hombre)
	ejerinterfaces.HumanoRespirando(Pedro)

	Maria := new(modelos.Mujer)
	ejerinterfaces.HumanoRespirando(Maria)*/

	//******Ejercicio Defer y Panic*********
	/*deferpanic.VemosDefer()
	deferpanic.EjemploPanic()*/

	canalUno := make(chan bool)
	go goroutines.MiNombreLento("Felipe Pabon", canalUno)

	fmt.Println("Estoy aqui")
	/*var x string
	fmt.Scanln(&x)*/

	//_ = <-canalUno //await _ no se captura el valor, solo se espera a que termine
	defer func() { canalUno <- true }() //defer se ejecuta al final de la funcion

}
