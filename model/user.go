package model

import (
	"time"
)

//Creamos definiciones de estructuras.

// User es una estructura que representa un usuario.
type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// AddUser es una funcion que agrega un usuario.
// Se debe agregar un * antes del tipo de dato para que sea un puntero.
func (user *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	user.Id = id
	user.Name = name
	user.CreatedAt = createdAt
	user.Status = status
}
