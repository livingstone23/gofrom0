package ejer_interfaces

import (
	"fmt"

	"github.com/livingstone23/gofrom0/interfaces"
)

func HumanRespirar(hu interfaces.Human) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
