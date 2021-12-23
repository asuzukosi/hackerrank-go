package main

import "fmt"

func main() {
	var len int
	var word string
	var key int

	fmt.Scanln(&len)
	fmt.Scanln(&word)
	fmt.Scanln(&key)
	fmt.Println(word)
	fmt.Println(CeaserCipher(word, key))
}
