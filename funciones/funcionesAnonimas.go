package funciones

import "fmt"

func Calculos() {
	//funciones anónimas
	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(suma(5, 6))

}
