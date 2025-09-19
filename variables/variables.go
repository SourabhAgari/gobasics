package variables

import "fmt"

var c, python, java bool

func Variables() {
	var a int = 20
	j := 30 // shorthand for variable declaration and initialization
	var l float64
	fmt.Println("Value of a is:", a)
	fmt.Println("Value of j is:", j)
	fmt.Println("Value of Global Varfiables c is:", c, "Python is:", python, "Java is:", java)
	fmt.Println("Value of l is:", l)
}
