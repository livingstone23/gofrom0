package iterations

import (
	"fmt"
)

func Iterate() {

	i := 0
	for i < 10 {
		fmt.Println("El valor de i es: ", i)
		i++
	}

}

func IterateSmall() {
	for i := 0; i < 10; i++ {
		fmt.Println("El valor de i es: ", i)
	}
}

// Iteracciones con saldos
func IterateWithJamps() {
	for i := 0; i <= 100; i += 5 {
		fmt.Println("El valor de i es: ", i)
	}
}

// Iteracciones con saldos
func IterateBreak() {
	for i := 10; i > 1; i-- {

		//Con el brak aborta
		if i == 3 {
			break
		}

		//Con el continue salta
		if i == 7 {
			continue
		}

		fmt.Println("El valor de i es: ", i)

	}
}
