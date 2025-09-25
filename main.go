package main

import (
	"fmt"

	"github.com/SourabhAgari/gobasics/arrays"
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

	// 8. Arrays in GoLang
	fmt.Println("Arrays Demonstration")
	fmt.Println("--------------------")
	arrays.ArrayFunctions()
	fmt.Println("--------------------")
	fmt.Println("End of Arrays Demonstration")

	// 9. Slices in GoLang
	fmt.Println("Slices Demonstration")
	fmt.Println("--------------------")
	references.CallSliceFunctions()
	fmt.Println("--------------------")
	fmt.Println("End of Slices Demonstration")

	// 10. Maps in GoLang
	fmt.Println("Maps Demonstration")
	fmt.Println("--------------------")
	references.CallMapFunctions()
	fmt.Println("--------------------")
	fmt.Println("End of Maps Demonstration")

}
