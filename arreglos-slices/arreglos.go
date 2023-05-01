package arreglosslices

import "fmt"

var tabla []int                            //slice
var tabla2 [10]int = [10]int{10, 0, 0, 59} //arreglo
var matriz [5][7]int

func MuestroArreglos() {
	tabla := [10]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	// for i := 0; i < len(tabla); i++ {
	// 	fmt.Println(tabla[i])
	// }
	fmt.Println(tabla)
	tabla[7] = 99
	fmt.Println(tabla)

	fmt.Println(tabla2)

	tabla3 := [10]string{"Colombia", "Peru", "Argentina", "Venezuela", "Ecuador", "Chile", "Bolivia", "Paraguay", "Uruguay", "Brasil"}
	fmt.Println(tabla3)

	for i := 0; i < len(tabla3); i++ {
		if i == 5 {
			fmt.Println(tabla3[i])
		}
	}

	matriz[3][5] = 76
	fmt.Println(matriz)

}
