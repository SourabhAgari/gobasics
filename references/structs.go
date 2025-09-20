package references

import "fmt"

type Vertex struct {
	X int
	Y int
}

func basicStruct() {
	v := Vertex{1, 2}
	a := &v
	a.X = 4
	println("value of a is :", a)
	fmt.Println("value of x is :", v.X, "Value of y is :", v.Y)

	v2 := Vertex{X: 5} // Y is 0 by default
	fmt.Println("value of v2 is :", v2)

	v3 := Vertex{}
	fmt.Println("value of v3 is :", v3)

	p := &Vertex{1, 2}
	fmt.Println("value of p is :", p.X)

}
func StructFunctions() {
	// call to basic Struct
	basicStruct()
}
