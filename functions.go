// MARK: - Part 4: Types
// Reference: https://golangbot.com/types/

package main

import "fmt"

func main() {
	var bill = calculateBill(200, 2)
	fmt.Println("My bill is", bill)

	// area, perimeter := getRectAreaAndPerimeter(10, 10)
	// fmt.Println("area:", area, "perimeter:", perimeter)

	area2, _, _ := getRectAreaAndPerimeter(12.4, 12.5)
	fmt.Println("area2:", area2)
}

func calculateBill(price int, amount int) int {
	var totalPrice = price * amount
	return totalPrice
}

// func getRectAreaAndPerimeter(length int, width int) (int, int) {
// 	var area = length * width
// 	var perimeter = (length * 2) + (width * 2)
// 	return area, perimeter
// }

func getRectAreaAndPerimeter(length float64, width float64) (area float64, perimeter float64, rectName string) {
	area = length * width
	perimeter = (2 * length) + (2 * width)
	return area, perimeter, rectName
}
