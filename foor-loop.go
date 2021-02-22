/* Foor Loop in Go*/

package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("--- ---")

	// for dengan statement
	for indx := 1; indx <= 5; indx++ {
		fmt.Println(indx)
	}

	// for manual for slice
	names := []string{"Made", "Erwin", "Ardiantha"}
	fmt.Println("--- ---")
	for j := 0; j < len(names); j++ {
		fmt.Println(names[j])
	}

	// for range
	slices := []string{"Ayu", "Indira", "Savitri"}
	for _, value := range slices {
		fmt.Println(value)
	}

}
