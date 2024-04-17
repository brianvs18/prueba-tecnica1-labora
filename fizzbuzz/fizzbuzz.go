package fizzbuzz

import "fmt"

func calculate(n int, multiplier int) bool {
	return n%multiplier == 0
}

func Execute(n int) {
	var store []string
	for number := 1; number <= n; number++ {
		switch true {
		case calculate(number, 15):
			store = append(store, "FizzBuzz")
		case calculate(number, 3):
			store = append(store, "Fizz")
		case calculate(number, 5):
			store = append(store, "Buzz")
		default:
			store = append(store, fmt.Sprintf("%v", number))
		}
	}
	fmt.Println(store)
}
