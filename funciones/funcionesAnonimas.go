package funciones

import "fmt"

//PAQUETE PARA APRENDER FUNCIONES ANONIMAS

func Calculos() {

	var numero3 int = 30
	var numero4 int = 243

	//Asignamos una funcion a una variable
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3 * numero4
	}

	fmt.Println(calculo(10, 25))

}
