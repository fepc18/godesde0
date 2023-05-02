package users

import (
	"fmt"
	"time"

	"github.com/fepc18/godesde0/modelos"
)

func AltaUsuario() {
	usuario := new(modelos.User)
	fmt.Println(usuario)
	usuario.AddUser(1, "Fernando", time.Now(), true)
	fmt.Println(usuario)
}
