package main

import "fmt"

func main() {
	testScoreGrade5 := 80
	testScoreGrade4 := 60
	testScoreGrade3 := 50
	testScore := -20

	hasGas := true
	hasKeyignition := true

	hasBurger := true
	hasSandwich := false

	if testScore >= testScoreGrade5 {
		fmt.Println("Top Mark")
	} else if testScore >= testScoreGrade4 {
		fmt.Println("Pass with distinction")
	} else if testScore >= testScoreGrade3 {
		fmt.Println("Pass with distinction")
	} else if testScore >= 0 {
		fmt.Println("Failed")
	} else {
		fmt.Println("Invalid test score")
	}

	if hasGas && hasKeyignition {
		fmt.Println("Can drive car")
	}

	if hasBurger || hasSandwich {
		fmt.Println("Can eat")
	}

	if !hasSandwich {
		fmt.Println("No sandwiches, so i shall starveeee")
	}
}
