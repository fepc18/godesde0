package iteraciones

import (
	"fmt"
)

func Iterar() {
	// While no existe en Go
	// for {
	// 	//break // break es para salir del for
	// 	//continue // continue es para saltar a la siguiente iteracion
	// }
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	//En una sola linea iterando de 5 en 5
	for i := 0; i < 10; i += 5 {
		fmt.Println(i)
	}
	fmt.Println("**iterando de 10 a 1, usando continue para saltar el 5**")
	for i := 10; i > 1; i-- {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
