package main

import (
	"fmt"
	"os"
	"prueba-tecnica1-labora/bug-tracking-system/presenters"
	"prueba-tecnica1-labora/fizzbuzz"
	menu_packaging "prueba-tecnica1-labora/menu-packaging"
	soccer_players "prueba-tecnica1-labora/soccer-players"
	"time"
)

func main() {
	for {
		var choice int
		fmt.Println(":: MENU PRINCIPAL :: ")
		fmt.Println("1. Ejercicio jugadores de futbol")
		fmt.Println("2. Ejercicio FizzBuzz")
		fmt.Println("3. Ejercicio Menu Packaging")
		fmt.Println("4. Ejercicio entregable")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			soccer_players.Execute()
			sleepOneSecond()
		case 2:
			fizzbuzz.Execute(15)
			sleepOneSecond()
		case 3:
			menu_packaging.Execute()
			sleepOneSecond()
		case 4:
			presenters.InitMenu()
			sleepOneSecond()
		case 5:
			fmt.Println("Hasta luego")
			sleepOneSecond()
			os.Exit(0)
		default:
			fmt.Println("Opción incorrecta, por favor intentelo de nuevo")
			sleepOneSecond()
		}
	}
}

func sleepOneSecond() {
	time.Sleep(1 * time.Second)
}
