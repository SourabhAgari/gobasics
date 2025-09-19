package hello

import (
	"fmt"
	"math/rand"
)

// if it is used by other packages,
// the function name should start with a capital letter
func SayHello() {
	fmt.Println("Hello, Go!")
}

func GenerateRandomNumber() {
	num := rand.Intn(100) // generates a random number between 0 and 99
	fmt.Println("Generating a random number...", num)
}
