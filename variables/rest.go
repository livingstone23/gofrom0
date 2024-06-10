package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestVariables() {
	//Definir variables in declarative way
	Nombre = "Livingstone"
	Estado = true
	Sueldo = 2350.66
	Fecha = time.Now()

	fmt.Println("Nombre =", Nombre)
	fmt.Println("Estado =", Estado)
	fmt.Println("Sueldo =", Sueldo)
	fmt.Println("Fecha =", Fecha)

}

func CastToText(number int) (bool, string) {
	//var text string
	text := strconv.Itoa(number)
	return true, text
}
