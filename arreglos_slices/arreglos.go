package arreglos_slices

import "fmt"

//arreglo, recordar que existe la posicion 0
var tabla [10]int

var tabla2 [10]int = [10]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

var tabla3 [10]string

//Creacion de una matriz
var matriz [5][10]int

func MuestraArreglos() {
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	fmt.Println(tabla2)

	tabla3[0] = "Livingstone"
	tabla3[1] = "Miguel"
	fmt.Println(tabla3)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[3][5] = 15
	fmt.Println(matriz)

}
