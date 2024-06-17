package funciones

import "fmt"

//FUNCIONES RECURSIVAS
func Exponencia(valor int) {
	if valor > 10000 {
		return
	}

	fmt.Println(valor)
	Exponencia(valor * 2)

}
