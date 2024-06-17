package funciones

import "fmt"

///DEVUELVE UNA funciona anonima que retorna un entero
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

//Permite realizar ofuscacion de codigo
func LLamarClosure() {
	tablaDel := 2
	MiTabla := tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
