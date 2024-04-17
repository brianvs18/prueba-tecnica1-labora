package presenters

import (
	"fmt"
	"prueba-tecnica1-labora/bug-tracking-system/auth"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
)

func InitMenu() {
	for {
		fmt.Println(":: EJERCICIO ENTREGABLE :: ")
		fmt.Println("Bienvenido/a")
		fmt.Println("1. Iniciar sesion")
		fmt.Println("0. Volver al menu principal")
		fmt.Print("Seleccione una opcion: ")
		choice := shared.GetScanner()

		switch choice {
		case 1:
		case 0:
			fmt.Println("Hasta luego")
			return
		default:
			fmt.Println("Opcion incorrecta, intente de nuevo")
		}
		fmt.Println(auth.Users)
	}
}
