package shared

import "fmt"

func GetScanner() int {
	var choice int
	fmt.Scanln(&choice)
	return choice
}
