package files

import (
	"fmt"
	"os"

	"github.com/fepc18/godesde0/ejercicios"
)

const fileName = "./files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.Multiplicar(8)
	//Crear un archivo
	// Si el archivo existe, lo sobreescribe
	archivo, err := os.Create(fileName)
	// Si hay un error al crear el archivo, imprime el error y termina la ejecución
	if err != nil {
		fmt.Println(err)
	}
	archivo.WriteString(texto)
	archivo.Close()

}

func SumaTabla() {
	var texto string = ejercicios.Multiplicar(3)
	//Crear un archivo
	if !Append(fileName, texto) {
		fmt.Println("Error al escribir el archivo")
	}

}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println(err)
		return false
	}
	arch.Close()
	return true

}
