package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/SourabhAgari/gobasics/variables"
)

func main() {
	// 1.hello world printing example
	// hello.SayHello()

	// 2.generate a random number
	//hello.GenerateRandomNumber()

	// 3. functions example

	// 4. variables
	variables.Variables()

	name := "Sourabh"
	fmt.Println(&name)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&name))
	fmt.Println(stringHeader.Data)
	fmt.Println(stringHeader.Len)

}
