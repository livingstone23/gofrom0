package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func GenerateTable() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el número para generar la tabla: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(numero1, " x ", i, " = ", numero1*i)
	}

}
