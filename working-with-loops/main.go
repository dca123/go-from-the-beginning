package main

import "fmt"

func main3() {
	arr := []string{"arg1", "arg2", "arg3"}
	for i, s := range arr {
		fmt.Printf("index: %d, item: %s \n", i, s)
	}
}

func main() {
	keepGoing := true
	answer := ""
	for keepGoing {
		fmt.Println("Type command: (quit to exit)")
		fmt.Scan(&answer)
		if answer == "quit" {
			break
		} else if answer == "print" {
			fmt.Println("Printing file")
		}
	}
	fmt.Println("program exit")
}
