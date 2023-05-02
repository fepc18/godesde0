package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Estoy en el middleware")

	result := operacionMidd(sumar)(5, 3)
	fmt.Println(result)

	result = operacionMidd(restar)(5, 3)
	fmt.Println(result)

	result = operacionMidd(multiplicar)(5, 3)
	fmt.Println(result)
}

func operacionMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
