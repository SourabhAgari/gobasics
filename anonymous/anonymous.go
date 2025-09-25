package anonymous

import "fmt"

func fibonnaci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonnaci(n-1) + fibonnaci(n-2)
}

func basicAnonymousFunc() {
	// these flavours can be called only once
	println("this is a basic anonymous function")
	func() {
		fmt.Println("2 * 8", 2*8)
	}()
	fmt.Println("end of basic anonymous functions")

	fmt.Println("anonymous function with parameters")
	func(a int, b int) {
		fmt.Println("a + b", a+b)
	}(3, 4)
}

func functionAsAVariable() {
	vars := func(a int, b int) int {
		return a + b
	}
	result := vars(3, 4)
	fmt.Println("result from function as a variable", result)

	// we can also assig a named function to a variable
	// but we cannot use a named function inside a named function in go

	fib := fibonnaci
	fmt.Println("fibonnaci of 7 is", fib(7))

}

// higher order function - function which takes another function as a parameter
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}
func functionAsAParam() {
	sum := applyOperation(3, 4, func(a int, b int) int {
		return a + b
	})
	fmt.Println("sum is", sum)

	multiply := applyOperation(3, 4, func(a int, b int) int {
		return a * b
	})
	fmt.Println("multiply is", multiply)
}

func CallOthers() {
	basicAnonymousFunc()
	functionAsAVariable()
	functionAsAParam()
}
