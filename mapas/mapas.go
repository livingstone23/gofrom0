package mapas

import (
	"fmt"
)

func MostrarMapas() {
	paises := make(map[string]string)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	//fmt.Println(paises)

	//fmt de valores determinados
	//fmt.Println(paises["Mexico"])

	//Definir un mapa con valores determinados
	paises2 := make(map[string]string, 5)

	paises2["Nicaragua"] = "Managua"
	paises2["Costa Rica"] = "San José"
	paises2["Honduras"] = "Tegucigalpa"
	paises2["El Salvador"] = "San Salvador"
	//fmt.Println(paises2)

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}
	//fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	}

	//Eliminamos un registro.
	delete(campeonato, "Chivas")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Real Madrid"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

}
