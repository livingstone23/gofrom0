package arreglos_slices

import "fmt"

var tablas []int = []int{2, 5, 4}

var arreglo [10]int = [10]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

func MuestroSlices() {

	fmt.Println(tablas)

	//Slices desdes la posicion 3 hasta el final
	porcion := arreglo[3:]
	fmt.Println(porcion)

	//Slices desde la posicion 0 hasta la 5
	porcion2 := arreglo[:5]
	fmt.Println(porcion2)

	//Slices desde la posicion 3 hasta la 6
	porcion3 := arreglo[3:6]
	fmt.Println(porcion3)

}

func Capicidad() {

	//Crear un slice con 20 elementos de capacidad inicial
	//Tendra 5 elementos y 20 de capacidad
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo: %d, Capacidad: %d\n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo: %d, Capacidad: %d\n", len(nums), cap(nums))

}
