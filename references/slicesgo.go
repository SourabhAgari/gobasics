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

func appendDemo() {
	var s []int
	fmt.Println(s)

	fmt.Println("Appending 0")
	s = append(s, 0)
	fmt.Println(s)

	fmt.Println("Appending 1")
	s = append(s, 1)
	fmt.Println(s)

	fmt.Println("Appending 2,3,4,5")
	// this shows that append is a variadic fuction
	s = append(s, 2, 3, 4, 5)
	fmt.Println(s)

	fmt.Println("Appending another slice")
	s = append(s, []int{6, 7, 8}...)
	fmt.Println(s)

}

func twoDSlice() {
	a := [][]string{
		{"a", "b"},
		{"c", "d"},
	}
	fmt.Println("2D Slice", a)

	// below one is not allowed in go lang as lang is statically typed.
	// b := [][]{
	// 	[]int{1,2,3},
	// 	[]string{"s","o","u"}
	// }
}

func loopOverSlices() {
	fmt.Println("Looping over slices")
	fmt.Println("using traditional for loop")
	s := []int{2, 3, 5, 7, 11, 13}
	for i := 0; i < len(s); i++ {
		fmt.Println("Element at index", i, "is", s[i])
	}

	fmt.Println("using range based for loop")
	for index, value := range s {
		fmt.Println("Element at index", index, "is", value)
	}
}

func CallSliceFunctions() {
	basicSlice()
	advancedSlice()
	extendingSliceLength()
	makeDemo()
	twoDSlice()
	appendDemo()
	loopOverSlices()
}
