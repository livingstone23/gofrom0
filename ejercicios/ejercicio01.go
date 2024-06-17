package ejercicios

import "strconv"

func ConvNumeric(text string) (int, string) {
	num, err := strconv.Atoi(text)

	//Manejo de errores
	if err != nil {
		return 0, "Hubo un error al convertir el texto a numero" + err.Error()

	}

	if num > 100 {
		return num, "El numero es mayor a 100"
	} else {
		return num, "El numero es menor o igual a 100"
	}

}
