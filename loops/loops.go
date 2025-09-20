package loops

import (
	"fmt"
	"math"
)

func forLoop() {
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// Counting backwards
	fmt.Println("\nCountdown:")
	for i := 5; i >= 1; i-- {
		fmt.Printf("T-minus %d\n", i)
	}
	fmt.Println("Blast off! 🚀")

	// Skip numbers (increment by 2)
	fmt.Println("\nEven numbers from 2 to 10:")
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func whileLoop() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("Sum is:", sum)
}

func pow(x, n, lin float64) float64 {
	v := math.Pow(x, n)
	if v < lin {
		return v
	}
	return lin
}

func rangeLoops() {
	fmt.Println("Demo Range loop over a slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	fmt.Println("********************")
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	fmt.Println("********************")
	for index := range fruits {
		fmt.Printf("Index: %d\n", index)
	}
	fmt.Println("********************")
	fmt.Println("--------------------")
}

func nestedLoops() {
	fmt.Println("Demo of nested loops:")
	for i := 1; i < 10; i++ {
		for j := 1; j < 5; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}
}

func CallOthers() {

	// demo of for loop
	fmt.Println("Demo of for loop:")
	forLoop()
	fmt.Println("--------------------")
	// demo of while loop
	fmt.Println("\nDemo of while loop:")
	whileLoop()
	fmt.Println("--------------------")
	// demo of expressions inside for loop
	fmt.Println("\nDemo of expressions inside for loop:")
	res := pow(3, 2, 10)
	fmt.Println(res)
	fmt.Println("--------------------")

	rangeLoops()
	nestedLoops()
}
