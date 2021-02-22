/*If Expression*/

package main

import "fmt"

func main() {
	name := "iin"

	if name == "rrwin" {
		fmt.Println("Hello rrwin")
	} else if name == "iin" {
		fmt.Println("Hello Darling😂")
	} else {
		fmt.Println("You're not rrwin")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("More than 5")
	} else {
		fmt.Println("Less than 5")
	}
}
