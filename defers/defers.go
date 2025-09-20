package defers

import "fmt"

// A defer statement in Go delays the execution of a function
// until the surrounding function returns.
// It's like a "cleanup" mechanism that ensures certain code runs before the function exits.

//Key points about defer:

//Defers are executed in LIFO (Last In, First Out) order
//Defers are executed even if the function panics
//Arguments to deferred functions are evaluated immediately
//Commonly used for cleanup operations like closing files or network connections

// basic defers
func basicDefer() {
	fmt.Println("Start of basicDefer function")
	defer fmt.Println("This is a deferred statement 1")
	fmt.Println("Middle of basicDefer function")
	fmt.Println("End of basicDefer function")
}

// multiple defers in go lang
func multipleDefers() {
	fmt.Println("Start of multipleDefers function")
	// executed in LIFO order
	defer fmt.Println("This is deferred statement 1")
	defer fmt.Println("This is deferred statement 2")
	defer fmt.Println("This is deferred statement 3")
	fmt.Println("End of multipleDefers function")
}

// function that invokes call to other functions in this package
func CallOthers() {
	basicDefer()
	multipleDefers()
}
