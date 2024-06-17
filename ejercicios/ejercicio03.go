package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero2 int
var err2 error
var texto string

func GenerateTable2() string {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Ingrese el número para generar la tabla: ")
		if scanner.Scan() {
			numero2, err2 = strconv.Atoi(scanner.Text())
			if err2 != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero2, i, numero2*i)
	}

	return texto

}
