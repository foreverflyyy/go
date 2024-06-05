package main

import "fmt"

func main() {
	numbers := []string{"1", "2", "3", "4", "5"}
	fmt.Println("Hello Go!")
	fmt.Println(numbers)

	if numbers[0] == "1" {
		fmt.Println("wow")
	}

	users := []string{"Tom", "Bob"}
	for key, value := range users {
		fmt.Println(key, value)
	}
}

func add(first, second int) int {
	return 10 + first + second
}
