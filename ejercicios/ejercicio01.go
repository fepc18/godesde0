package ejercicios

import "strconv"

func Funcion(valor string) (int, string) {

	resString := ""
	resInt, err := strconv.Atoi(valor)
	if err != nil {
		if resInt > 100 {
			resString = "Mayor a 100"
		} else {
			resString = "Menor o igual a 100"
		}
	}
	return resInt, resString

}
