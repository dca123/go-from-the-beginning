package main

import "fmt"

type Address struct {
	city   string
	street string
	postal string
}

type Person struct {
	name    string
	address Address
}

type Employee struct {
	Address
	company string
}

func (e Employee) string() string {
	return fmt.Sprintf("Company: %s", e.company)
}

func main() {
	person := Person{
		name: "Devinda",
		address: Address{
			city:   "Colombo",
			street: "idk",
			postal: "111222",
		},
	}
	fmt.Println(person)

	employee := Employee{
		Address: Address{
			city:   "LA",
			postal: "10101010",
		},
		company: "Hello World",
	}

	fmt.Println(employee.string())
	fmt.Println(employee.company)
}
