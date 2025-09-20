package main

import (
	"fmt"

	"github.com/SourabhAgari/gobasics/defers"
	"github.com/SourabhAgari/gobasics/loops"
	"github.com/SourabhAgari/gobasics/references"
	"github.com/SourabhAgari/gobasics/switchcase"
)

func main() {
	// 1.hello world printing example
	// hello.SayHello()

	// 2.generate a random number
	//hello.GenerateRandomNumber()

	// 3. functions example

	// 4. variables
	//variables.Variables()

	// 5. Loops
	loops.CallOthers()

	// 6. Switch Case
	switchcase.CallOthers()

	//7. Defers in GoLang
	fmt.Println("Defers Demonstration")
	fmt.Println("--------------------")
	defers.CallOthers()
	fmt.Println("--------------------")
	fmt.Println("End of Defers Demonstration")

	//7. Pointers in GoLang
	fmt.Println("Pointers Demonstration")
	fmt.Println("--------------------")
	references.PointersFunctions()
	fmt.Println("--------------------")
	fmt.Println("End of Pointers Demonstration")

	// 8. Structs in GoLang
	fmt.Println("Structs Demonstration")
	fmt.Println("--------------------")
	references.StructFunctions()
	fmt.Println("--------------------")
	fmt.Println("End of Structs Demonstration")

	sum := 0
	i := 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println("Sum is:", sum)
}
