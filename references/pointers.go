package references

import (
	"fmt"
)

// pointers
// pointer is just a variable that holds the memory address of another variable
// & operator gives the memory address of a variable
// * operator gives the value at the address (also can be called as dereferencing operator)

func pointersDemo() {
	// define variables
	name := "Sourabh"

	// contains the address of the name variable
	addrName := &name

	// update the existing value in memory so no copy
	*addrName = "Ragari"
	println(name)

	fmt.Println("pointers example 2")
	j := 2701
	fmt.Println("Value of j before is:", j)
	p := &j
	*p = *p / 37
	fmt.Println("Value of j after is:", j)

}

func PointersFunctions() {
	pointersDemo()
}
