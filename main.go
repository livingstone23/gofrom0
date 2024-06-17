package main

import (
	d "github.com/livingstone23/gofrom0/defer_panic"
)

// main function
func main() {

	//Primera funcion declarada
	//variables.ShowIntegers()

	//Segunda funcion declarada
	//variables.RestVariables()

	//Tercera funcion declarada
	//estado, text := variables.CastToText(10)
	//fmt.Println("Estado: ", estado, " Text: ", text)

	/*
		if os := runtime.GOOS; os == "linux" || os == "OS X." {
			fmt.Println("Esto no es Windows, es: ", os)
		} else {
			fmt.Println("Esto es Windows")
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "OS X":
			fmt.Println("Esto es OS X")
		default:
			fmt.Printf("%s \n", os)
		}
	*/

	//Ejercicio 01
	//number, text := ejercicios.ConvNumeric("500")
	//fmt.Println("Numero: ", number, " Texto: ", text)

	//Mostar y aceptar datos en GO
	//keyboard.IncomeNumber()

	//Ciclos en GO
	//iterations.Iterate()
	//iterations.IterateSmall()
	//iterations.IterateWithJamps()
	//iterations.IterateBreak()

	//Ejercicio 02
	//ejercicios.GenerateTable()

	//fmt.Println(ejercicios.GenerateTable2())

	//"github.com/livingstone23/gofrom0/files"
	//Manejo de archivos en GO.
	//files.SaveTable()
	//files.AddTable()
	//files.ReadFile()

	//"github.com/livingstone23/gofrom0/funciones"
	//funciones.Calculos()
	//funciones.LLamarClosure()
	//funciones.Exponencia(2)

	//"github.com/livingstone23/gofrom0/arreglos_slices"
	//Arreglos y slices
	//arreglos_slices.MuestraArreglos()
	//arreglos_slices.MuestroSlices()
	//arreglos_slices.Capicidad()

	//"github.com/livingstone23/gofrom0/mapas"
	//mapas.MostrarMapas()

	//github.com/livingstone23/gofrom0/users
	//users.AltaUsuario()

	//e "github.com/livingstone23/gofrom0/ejer_interfaces"
	//m "github.com/livingstone23/gofrom0/model"
	//Livingstone := new(m.Men)
	//e.HumanRespirar(Livingstone)

	//Maria := new(m.Woman)
	//e.HumanRespirar(Maria)

	//d.VemosDefer()
	d.EjemploPanic()

}
