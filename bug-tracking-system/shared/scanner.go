package shared

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetScanner() int {
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func GetScannerString() string {
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	value = strings.TrimSpace(value)
	return value
}
