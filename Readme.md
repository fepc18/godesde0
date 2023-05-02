1. archivo. Go.mod // for Packages
   go mod init github.com/fepc18/godesde0

2. fmt --> para imprimir en consola

3. execute --> go run main.go

4. Build --> go build main.go
   ejecutarlo ./main o ./main.exe

5. Syntaxis: Mayuscula Publica (minuscula local)
   Ej: MuestroEnteros(){
6. Declaracion de variables
   // declarative
   var x int
   x = 5
   // infered
   y := 10

7. Los archivos dentro del mismo paquete pueden acceder a los metodos y variables publicas de los otros paquetes sin tener q importarlos.

8. Parse string to int --> strconv.Atoi(valor)
   Parse int to string --> strconv.Itoa(numero)

9. Files

10. Funciones Anonimas y Closures
    //Funciones anonimas, funciones que se pueden enviar como parametro
    //Closures: Funciones que se asignan a una variable y luego se ejecutan
    //Recursiuon

11. Diferencia array y slice: Slice no se define las posiciones iniciales
    Se recorren con for
    para asignar un valor
    tabla[4]=3
    var tabla []int //slice
    var tabla2 [10]int = [10]int{10, 0, 0, 59} //arreglo
    var matriz [5][7]int

    make para crear slices con capacidad definida, buen uso para el rendimiento

12. Map
    Llave Valor
    paises := make(map[string]string) //5 es la capacidad
    for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	}
   Agregar campeonato["Atletico"] = 33 //Si no existe lo agrega
   delete(campeonato, "Real Madrid")

13. Estructuras, no existe clases, sino Type
    Para agregar metodos se crea por fuera y se debe referenciar el type (*)
    Cuando se crea, no usa nulos sino asigna valores por defecto

14. Interfaces: Modelo de comportamiento 
    Con que struct tenga implementados, se infiera que esta implementando una interfaz

    Herencia::::
    type Mujer struct {
         Hombre
   }
15. Defer: Funcion al ejecutarse al salir de una funcion
    Se puede tener varios defer
    Panic:   abortar la ejecucion con un mensaje
    Recover: capturar cualquier panic, ej. capturarlos y guardar en un lock antes   de   abortar un sistema --> debe usarse con el defer
