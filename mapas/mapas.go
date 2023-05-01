package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string) //5 es la capacidad
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	}
	delete(campeonato, "Real Madrid")
	fmt.Println("****Despues de eliminar Real Madrid****")
	fmt.Println(campeonato)

	campeonato["Atletico"] = 33 //Si no existe lo agrega
	fmt.Println("****Despues de agregar****")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

}
