package main

import "fmt"

func main() {
	fruits := [...]string{
		"buah naga", "anggur", "ceri", "apel", "pepaya",
	}

	for _, buah := range fruits {
		fmt.Println(buah)
	}
}
