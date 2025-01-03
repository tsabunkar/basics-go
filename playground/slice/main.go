package main

import "fmt"

func main() {
    fmt.Println("Lets learn about slice make func")
	var s []byte
	s = make([]byte, 5, 5) // func make([]T, len, cap) []T
	fmt.Println(s) // [0 0 0 0 0]
	
	fmt.Println("capacity argument is omitted")
	somitted := make([]byte, 5)
	fmt.Println(somitted)

}
