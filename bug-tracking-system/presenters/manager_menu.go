package presenters

import (
	"fmt"
	"prueba-tecnica1-labora/bug-tracking-system/services/manager"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
)

func InitManagerMenu() {
	for {
		fmt.Println(":: Bienvenido MANAGER ::")
		fmt.Println("1. Crear tarea")
		fmt.Println("2. Generar reporte")
		fmt.Println("0. Cerrar sesion")
		fmt.Print("Seleccione una opcion: ")
		choice := shared.GetScanner()

		switch choice {
		case 1:
			fmt.Print("Ingrese el nombre de la tarea: ")
			name := shared.GetScannerString()
			manager.CreateTask(name)
		case 2:
		case 0:
			fmt.Println("Hasta pronto")
			return
		default:
			fmt.Println("Opcion incorrecta, intente nuevamente")
		}
	}
}
