package main

import "fmt"

func main() {
    fmt.Println("Lets learn about slice make func")
	var s []byte
	s = make([]byte, 5, 5) // func make([]T, len, cap) []T
	fmt.Println(s) // [0 0 0 0 0]
	
	fmt.Println("capacity argument is omitted")
	sOmitted := make([]byte, 5) // capacity default value is specified to length
	fmt.Println(sOmitted) // [0 0 0 0 0]

	fmt.Println(len(sOmitted)) // 5
	fmt.Println(cap(sOmitted)) // 5

}
