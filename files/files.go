package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/livingstone23/gofrom0/ejercicios"
)

func SaveTable() {
	var texto string = ejercicios.GenerateTable2()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo")
		return
	}

	//Escribir en el archivo
	fmt.Fprintln(file, texto)
	//El archivo que se abra debe cerrarse
	file.Close()
}

var fileName string = "./files/txt/tabla.txt"

func AddTable() {
	var texto string = ejercicios.GenerateTable2()

	//Reduce comparacion
	//if Append(fileName, texto) == false {

	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar el archivo")
	}
}

func Append(fileName string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}

	//El archivo que se abra debe cerrarse
	arch.Close()
	return true
}

func ReadFile() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		register := scanner.Text()
		fmt.Println("> " + register)
	}

	file.Close()
}
