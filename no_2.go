package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	fmt.Print("Tugas 2: Reverse Strings per Kata\n")
	fmt.Print("Masukkan 3 kata : ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	kata := strings.Fields(input)

	if len(kata) < 3 {
		fmt.Println("3 Kata minimal bro -_-")
		return
	}

	for i := 0; i < len(kata); i++ {
		kata[i] = reverseString(kata[i])
	}

	kebalikan := strings.Join(kata, " ")
	fmt.Println("Hasil Reversenya : ", kebalikan)
}

func reverseString(s string) string {
	balik := []rune(s)
	for i, j := 0, len(balik)-1; i < j; i, j = i+1, j-1 {
		balik[i], balik[j] = balik[j], balik[i]
	}
	return string(balik)
}
