package interfaces

// Human es una interfaz que representa a un humano.
//Se definen los metodos
type Human interface {
	Respirar()
	Pensar()
	Comer()
	Sexo() string
}
