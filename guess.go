package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please think of a number between 1 and 100")
	fmt.Println("Press ENTER when ready")
	scanner.Scan()
	
	low := 1
	high := 100


	for {
		guess := (low + high) / 2
		fmt.Println("I guess the number is ", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()
		response := scanner.Text()

		switch response {

			case "a":
				high = guess

			case "b":
				low = guess
			
			case "c":
				fmt.Println("I won!")
				return

			default:
				fmt.Println("Let me repeat")
		}
	}
}
