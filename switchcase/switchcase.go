package switchcase

import "runtime"

func switchDemo() {
	os := runtime.GOOS
	switch os {
	case "darwin":
		println("Mac OS")
	case "linux":
		println("Linux OS")
	default:
		println(os)
	}
}

func CallOthers() {
	// demo of switch case
	switchDemo()
}
