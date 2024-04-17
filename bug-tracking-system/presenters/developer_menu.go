package presenters

import (
	"fmt"
	"prueba-tecnica1-labora/bug-tracking-system/services/developer"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
)

func InitDeveloperMenu() {
	for {
		fmt.Println(":: Bienvenido DEVELOPER ::")
		fmt.Println("Lista de tareas pendientes")
		developer.GetAllPendingTasks()
		fmt.Println("1. Actualizar estado de la tarea")
		fmt.Println("0. Cerrar sesion")
		fmt.Print("Seleccione una opcion: ")
		choice := shared.GetScanner()

		switch choice {
		case 1:
			fmt.Print("Ingrese el ID de la tarea: ")
			id := shared.GetScannerString()
			fmt.Println("Ingrese uno de los siguientes estados")
			fmt.Print("1. PENDING - 2. RESOLVED: ")
			selection := shared.GetScanner()
			developer.UpdateTaskStatus(id, selection)
		case 0:
			fmt.Println("Hasta pronto")
			return
		default:
			fmt.Println("Opcion incorrecta, intente nuevamente")
		}
	}
}
