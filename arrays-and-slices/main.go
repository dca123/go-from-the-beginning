package main

import "fmt"

func main() {
	var logs []string
	var command string
	for {
		fmt.Print("command> ")
		fmt.Scan(&command)

		if command == "new" {
			var entry string
			fmt.Scan(&entry)
			logs = append(logs, entry)
		} else if command == "list" {
			for _, log := range logs {
				fmt.Println(log)
			}
		} else if command == "quit" {
			break
		} else {
			fmt.Println("Invalid command, valid commands are new, list & quit")
		}
	}
}
