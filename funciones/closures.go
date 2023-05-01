package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}

}

func LllamarClosure() {
	tablaDel2 := tabla(2) // tablaDel2 es una función que retorna otra función
	tablaDel3 := tabla(3) //se define

	for i := 0; i < 10; i++ {
		fmt.Println(tablaDel2()) //se llama a la función que retorna
		fmt.Println(tablaDel3())
	}
}
