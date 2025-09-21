package references

import "fmt"

type Person struct {
	Name string
	Age  int
}

func basicSlice() {
	// directly create a slice
	// go will automatically create a backing array
	s := []int{2, 3, 5, 7, 9, 11}
	fmt.Println("Initial Slice", s)

	fmt.Println("Appending to slice")
	s = append(s, 13)
	fmt.Println("After Appending elemens to slice", s)

}

func advancedSlice() {

	persons := []Person{
		{"Sourabh", 28},
		{"Sonal", 23},
	}

	fmt.Println("Initial Slice of Structs", persons)

	fmt.Println("Appending to slice of structs")
	persons = append(persons, Person{"NewPerson", 30})
	fmt.Println("After Appending elemens to slice of structs", persons)
}

func extendingSliceLength() {
	// this is possible due to the backing array
	// but be careful while doing this
	// as it can lead to unexpected results
	// if the backing array is not large enough
	// to accommodate the new length
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Initial Slice", s)

	s = s[:0]
	fmt.Println("After extending length by 1 steps", s)

	s = s[:4]
	fmt.Println("After extending length by 4 steps", s)

	fmt.Println("test")
	test := make([]int, 3, 5)
	test[0], test[1] = 1, 2
	fmt.Println("test slice", test)
}

func makeDemo() {
	a := make([]int, 5)
	fmt.Println("make demo slice", a)

	b := make([]int, 0, 5)
	fmt.Println("make demo slice with 0 length and 5 capacity", b)

	c := b[:2]
	fmt.Println("make demo slice with extended length", c)

	d := c[2:5]
	fmt.Println("make demo slice with extended length to full capacity", d)
}

func CallSliceFunctions() {
	basicSlice()
	advancedSlice()
	extendingSliceLength()
	makeDemo()
}
