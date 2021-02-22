/* Switch Expression */

package main

import "fmt"

func main() {
	name := "Erwin"

	switch name {
	case "Erwin":
		fmt.Println("Is Me")
	case "Iin":
		fmt.Println("Is Me")
	default:
		fmt.Println("*** ** *")
	}

	// short expression
	switch length := len(name); length < 5 {
	case true:
		fmt.Println("Less than 5")
	case false:
		fmt.Println("More than 5")
	}

	// switch tanpa tanda kondisi
	panjang := len(name)
	switch {
	case panjang > 5:
		fmt.Println("Less than 5")
	case panjang < 5:
		fmt.Println("More than 5")
	default:
		fmt.Println("Anying ilang gublug")
	}
}
