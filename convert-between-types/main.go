package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	arg2 := os.Args[2]

	num1, err1 := strconv.Atoi(arg)
	num2, err2 := strconv.Atoi(arg2)

	if err1 != nil || err2 != nil {
		panic("Error converting input, have you entered two numbers ? ")
	}

	result := add(num1, num2)
	fmt.Printf("Result is %d", result)
}

func add(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}
