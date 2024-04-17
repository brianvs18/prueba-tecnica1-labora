package soccer_players

import (
	"fmt"
	"strings"
)

func Execute() {
	// Lista de jugadores actuales
	players := []string{"John Smith", "David Johnson", "Michael Garcia", "Sarah Williams", "Daniel Martinez", "Emily Brown", "James Rodriguez", "Sophia Lee", "Lucas Oliveira", "Olivia Taylor", "Mateo Hernandez", "Emma Wilson", "Gabriel Silva"}

	// Registro de nuevos jugadores
	newPlayers := []string{"New Player 1", "New Player 2", "New Player 3", "New Player 4", "New Player 5", "New Player 6", "New Player 7"}

	// Eliminar los primeros 7 jugadores de la liga anterior
	players = players[7:]

	// Agregar nuevos jugadores
	for _, name := range newPlayers {
		players = append(players, name)
	}

	// Mostrar lista completa de jugadores
	fmt.Println("Lista completa de jugadores:")
	for i, player := range players {
		fmt.Printf("%d. %s\n", i+1001, sanitizeName(player))
	}
}

// Función para eliminar caracteres especiales y números del nombre del jugador
func sanitizeName(name string) string {
	// Eliminar caracteres especiales y números
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
		}
		return -1
	}, name)
}
