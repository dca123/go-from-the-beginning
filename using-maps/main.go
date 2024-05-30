package main

import "fmt"

func main() {
	var phonebook = map[string]string{"Alice": "555-123", "Bob": "555-124", "Jean": "555-125"}
	fmt.Println("Welcome to your phone book")
	for {
		fmt.Print("Command> ")
		var response string
		fmt.Scanln(&response)

		if response == "list" {
			for name, number := range phonebook {
				fmt.Printf("%s %s\n", name, number)
			}
		} else if response == "lookup" {
			fmt.Print("Enter name: ")
			var lookup string
			fmt.Scanln(&lookup)
			if value, exists := phonebook[lookup]; exists {
				fmt.Printf("%s has number %s\n", lookup, value)
			} else {
				fmt.Print("Contact doesn't exist, do you want to add it? y/n: ")
				var addContact string
				fmt.Scanln(&addContact)
				if addContact == "y" {
					fmt.Printf("Enter contact: %s ", lookup)
					var number string
					fmt.Scanf("%s", &number)
					phonebook[lookup] = number
				}
			}
		} else if response == "store" {
			fmt.Print("Enter contact: ")
			var name, number string
			fmt.Scanf("%s %s", &name, &number)
			phonebook[name] = number

		} else if response == "quit" {
			break
		} else {
			fmt.Println("Invalid command")
		}
	}
	fmt.Println("Bye !")

}
