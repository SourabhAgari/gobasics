package switchcase

import (
	"fmt"
	"runtime"
)

func switchDemo() {
	/**
	os := runtime.GOOS
	switch os {
	case "darwin":
		println("Mac OS")
	case "linux":
		println("Linux OS")
	default:
		println(os)
	}
		**/

	// can be written like this also
	switch os := runtime.GOOS; os {
	case "darwin": // can also be constants and expressions
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux OS")
	default:
		fmt.Println(os)
	}
}

func CallOthers() {

	// demo of switch case
	switchDemo()
}
