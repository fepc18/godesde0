package funciones

func Exponencial(numero int) int {
	if numero == 0 {
		return 1
	}
	return Exponencial(numero-1) * 2

}
