package variables

import "fmt"

func MuestroEnteros() {
	// declarative
	var x int
	x = 5
	fmt.Println("int declarative", x)

	// infered
	y := 10
	fmt.Println("int infered", y)

}
