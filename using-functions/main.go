package main

import (
	"fmt"
	"strconv"
)

func main() {
	log("Hello there !")
	log(strconv.Itoa(subtract(10, 5)))
}

func log(message string) {
	fmt.Println(message)
}

func subtract(num1 int, num2 int) (sum int) {
	sum = num1 - num2
	return
}
