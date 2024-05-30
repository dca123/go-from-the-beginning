package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

func Divide(nominator int, divider int) float32 {
	defer errorHandler()
	if divider == 0 {
		panic("cant divide by 0")
	}
	return float32(nominator) / float32(divider)
}

func errorHandler() {
	if r := recover(); r != nil {
		log.Println(r, string(debug.Stack()))
	}
}

func main() {
	f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)

	log.Println("Starting program")
	result := Divide(100, 0)
	fmt.Println(result)

	result = Divide(100, 10)
	fmt.Println(result)

	f.Close()
}
