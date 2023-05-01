package ejercicios

import "fmt"

func Multiplicar(number int) string {
	fmt.Println(":::::Tabla de multiplicar del", number)
	var texto string
	for i := 0; i <= 10; i++ {
		texto += fmt.Sprintln("Multiplicando", i, "por", number, "es igual a", i*number)
	}
	return texto

}
