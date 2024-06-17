package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Libreria bufio permite interactura con el teclado

var numero1 int
var numero2 int
var label string
var err error

func IncomeNumber() {
	scanner := bufio.NewScanner(os.Stdin)
	//os.Stdin es un archivo que representa la entrada estandar
	//en el caso de un sistema es el teclado
	fmt.Println("Ingrese número 1: ")
	if scanner.Scan() {
		//El texto ingresado se almacena en la variable text
		//Todo lo que entra por bufio entra como texto
		//Se convierte el texto a un numero
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingreasdo es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese número 2: ")
	if scanner.Scan() {
		//El texto ingresado se almacena en la variable text
		//Todo lo que entra por bufio entra como texto
		//Se convierte el texto a un numero
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingreasdo es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda ")
	if scanner.Scan() {
		//El texto ingresado se almacena en la variable text
		//Todo lo que entra por bufio entra como texto
		label = scanner.Text()
	}

	fmt.Println(label, numero1*numero2)

}
