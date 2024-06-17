package users

import (
	"fmt"
	"time"

	"github.com/livingstone23/gofrom0/model"
)

func AltaUsuario() {
	//Creamos un nuevo usuario
	u := new(model.User)
	u.AddUser(1, "Livingstone", time.Now(), true)

	fmt.Println("Usuario: ", u)

}
