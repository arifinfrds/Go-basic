// MARK: - Part 4: Types
// Reference: https://golangbot.com/types/

package main

import "fmt"

func main() {
	// MARK: - Type conversion
	var firstNumber = 10
	var secondNumber = 55.6
	var total = firstNumber + int(secondNumber)

	fmt.Println("total is", total)
}
