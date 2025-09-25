package references

import "fmt"

// defining map global variable
var m map[string]Vertex

func basicMapDemo() {
	/**
	// will initialize the map with some default values
	m = make(map[string]Vertex)
	fmt.Println(m)
	// assignig a value to a map datastructure
	m["Bell Labs"] = Vertex{40, -74}
	fmt.Println(m)
	**/
	// with syntactic sugar
	var m = map[string]Vertex{
		"Bell Labs": {40, -74},
		"Google":    {37, -122},
	}
	fmt.Println(m["Bell Labs"])
}

func deleteFunc() {
	m := make(map[string]string)

	m["name"] = "Sourabh"
	m["name2"] = "sonal"

	fmt.Println("map before deleting an element", m)

	delete(m, "name2")
	v, ok := m["name2"]
	fmt.Println("map after deleting an element", v, ok)

}

func loopOverMaps() {
	m := map[string]int{
		"Sourabh": 28,
		"Sonal":   23,
		"Ankit":   30,
	}

	for k, v := range m {
		fmt.Println("Key is :", k, "Value is :", v)
	}

	fmt.Println("Looping over only keys")
	for k := range m {
		fmt.Println("Key is :", k)
	}

	fmt.Println("Looping over only values")
	for _, v := range m {
		fmt.Println("Value is :", v)
	}
}

func CallMapFunctions() {
	basicMapDemo()
	deleteFunc()
	loopOverMaps()
}
