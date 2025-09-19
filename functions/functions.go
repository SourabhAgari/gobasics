package functions

import (
	"fmt"
)

// 1. normal function
func Greet() {
	fmt.Println("Hello, Go!")
}

// 2. function with parameters
func Add(a int, b int) {
	fmt.Println("Addition is:", a+b)
}

// 3. function with return values
func Multiply(a int, b int) int {
	return a * b
}

// 4. function with multiple return values
func DivideAndRemainder(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 5. Named Return Values
func Split(sum int) (x, y int) {
	x = sum * 9 / 10
	y = sum - x
	return
}

// 6. variadic functions
func SumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
