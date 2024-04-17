package menu_packaging

import (
	"fmt"
	"os"
	"time"
)

var (
	adminUserName      = "admin"
	adminPassword      = "root"
	supervisorUserName = "seeker223"
	supervisorPassword = "seekr"
	adminActions       = []string{"Crear laborer", "Eliminar laborer"}
	laborers           []string
	actionCounter      = make(map[string]int)
)

func Execute() {
	for {
		fmt.Println("Bienvenido a Labora")
		fmt.Println("1. Iniciar sesión como Administrador")
		fmt.Println("2. Iniciar sesión como Supervisor")
		fmt.Println("3. Salir")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			adminLogin()
		case 2:
			supervisorLogin()
		case 3:
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}

func adminLogin() {
	fmt.Println("Ingrese las credenciales de Administrador")
	fmt.Print("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == adminUserName && password == adminPassword {
		fmt.Println("Inicio de sesión exitoso como Administrador")
		adminMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo.")
	}
}

func supervisorLogin() {
	fmt.Println("Ingrese las credenciales de Supervisor")
	fmt.Print("Usuario: ")
	var username string
	fmt.Scanln(&username)
	fmt.Print("Contraseña: ")
	var password string
	fmt.Scanln(&password)

	if username == supervisorUserName && password == supervisorPassword {
		fmt.Println("Inicio de sesión exitoso como Supervisor")
		supervisorMenu()
	} else {
		fmt.Println("Credenciales incorrectas. Por favor, intente de nuevo.")
	}
}

func adminMenu() {
	for {
		fmt.Println("Menú de Administrador")
		for i, action := range adminActions {
			fmt.Printf("%d. %s\n", i+1, action)
		}
		fmt.Println("0. Cerrar sesión")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createLaborer()
		case 2:
			deleteLaborer()
		case 3:
			fmt.Println("Acción: Generar archivo de texto con mensajes personalizados")
			generateTextFile()
		case 0:
			fmt.Println("Cerrando sesión de Administrador")
			return
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}

func createLaborer() {
	laborers = append(laborers, fmt.Sprintf("laborer %d", len(laborers)+1))
	fmt.Println("Laborer creado exitosamente.")
	for i := 0; i < len(laborers); i++ {
		fmt.Println(laborers[i])
	}

}

func deleteLaborer() {
	if len(laborers) == 0 {
		fmt.Println("No hay laborers para eliminar.")
		return
	}

	laborers = laborers[:len(laborers)-1]
	fmt.Println("Laborer eliminado exitosamente.")
	for i := 0; i < len(laborers); i++ {
		fmt.Println(laborers[i])
	}
}

func supervisorMenu() {
	for {
		fmt.Println("Menú de Supervisor")
		fmt.Println("1. Crear registro de administrador")
		fmt.Println("0. Cerrar sesión")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Acción: Crear registro de administrador")
			createAdminRecord()
		case 0:
			fmt.Println("Cerrando sesión de Supervisor")
			return
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}

func generateTextFile() {
	fileName := fmt.Sprintf("mensajes_%s.txt", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear el archivo: %v\n", err)
		return
	}
	defer file.Close()

	// Escribir los mensajes en el archivo
	for _, laborer := range laborers {
		_, err := file.WriteString(fmt.Sprintf("%s\n", laborer))
		if err != nil {
			fmt.Printf("Error al escribir en el archivo: %v\n", err)
			return
		}
	}

	fmt.Printf("Archivo %s creado exitosamente.\n", fileName)
}

func createAdminRecord() {
	fileName := fmt.Sprintf("registro_admin_%s.txt", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear el archivo de registro: %v\n", err)
		return
	}
	defer file.Close()

	// Escribir los mensajes en el archivo
	for _, action := range adminActions {
		_, err := file.WriteString(fmt.Sprintf("%s: %d\n", action, actionCounter[action]))
		if err != nil {
			fmt.Printf("Error al escribir en el archivo de registro: %v\n", err)
			return
		}
	}

	fmt.Printf("Archivo de registro %s creado exitosamente.\n", fileName)
}
