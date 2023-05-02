package deferpanic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Primer Mensaje")

	defer fmt.Println("Mensaje Final")

	fmt.Println("Mensaje Segundo")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			fmt.Println("Recuperando el Panic")
			log.Fatalf("ocurrio un error que genero Panic  %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
