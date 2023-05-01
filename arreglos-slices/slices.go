package arreglosslices

import "fmt"

var tablaSlice []int = []int{3, 7, 9}                       //slice
var arreglo [10]int = [10]int{10, 3, 4, 59, 44, 22, 64, 78} //arreglo
func MuestroSlice() {
	fmt.Println(tablaSlice)

	porcion := arreglo[3:] //desde el indice 3 hasta el final
	fmt.Println(porcion)

	porcion2 := arreglo[:3] //desde el inicio hasta el indice 3
	fmt.Println(porcion2)

	porcion3 := arreglo[3:7] //desde el indice 3 hasta el indice 7
	fmt.Println(porcion3)

}

func Capacidades() {
	fmt.Println("Capacidad", cap(tablaSlice))
	fmt.Println("Capacidad", cap(arreglo))

	elementos := make([]int, 5, 20) //slice de 5 elementos y capacidad de 20
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
