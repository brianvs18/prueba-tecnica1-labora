package presenters

import (
	"fmt"
	"prueba-tecnica1-labora/bug-tracking-system/services/tester"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
)

func InitTesterMenu() {
	for {
		fmt.Println(":: Bienvenido TESTER ::")
		fmt.Println("Lista de tareas completadas")
		tester.GetAllCompletedTasks()
		fmt.Println("1. Crear bug")
		fmt.Println("2. Actualizar estado del bug")
		fmt.Println("3. Ver listado de bugs")
		fmt.Println("0. Cerrar sesion")
		fmt.Print("Seleccione una opcion: ")
		choice := shared.GetScanner()

		switch choice {
		case 1:
			fmt.Print("Ingrese una descripcion del bug: ")
			description := shared.GetScannerString()
			fmt.Print("Ingrese el ID de la tarea: ")
			taskId := shared.GetScannerString()
			tester.CreateBug(description, taskId)
		case 2:
			fmt.Println("Listado de bugs")
			tester.GetAllBugs()
			fmt.Print("Ingrese el ID del bug: ")
			id := shared.GetScannerString()
			fmt.Println("Ingrese uno de los siguientes estados")
			fmt.Print("1. PENDING - 2. RESOLVED: ")
			selection := shared.GetScanner()
			tester.UpdateBugStatus(id, selection)
		case 3:
			tester.GetAllBugs()
		case 0:
			fmt.Println("Hasta pronto")
			return
		default:
			fmt.Println("Opcion incorrecta, intente nuevamente")
		}
	}
}
