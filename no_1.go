package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Tugas 1: Hello Universe\n")
	fmt.Print("Inputkan Angka: ")
	fmt.Scan(&input)

	result := "Hello Universe!"
	
	if input != 42 {
		result = fmt.Sprintf("Input Angkamu adalah : %d", input)
	}
	fmt.Println(result)
}