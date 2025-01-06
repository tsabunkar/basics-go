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
	fmt.Println(cap(sOmitted)) // 
	
	// A slice can also be formed by “slicing” an existing slice or array.

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	
	fmt.Println(b) //[103 111 108 97 110 103]
	fmt.Println(b[2:]) //[108 97 110 103] <- creates a slice leaving first two elements
	fmt.Println(b[:2]) //[103 111] <- creates a slice of first two elements
	fmt.Println(b[:]) //[103 111 108 97 110 103]
	fmt.Println(b[1:4]) //[111 108 97] <- creates a slice including elements 1 through 3 of b 

	// creating slice from an array of string

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s1 := x[:] // a slice referencing the storage of x
	fmt.Println(s1)

	// Slicing does not copy the slice’s data. It creates a new slice value that points to the original array.
	// This makes slice operations as efficient as manipulating array indices.

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	fmt.Printf("%s \n", d) // Printing String for bytes value [road]
	fmt.Printf("%s \n",e)// [ad]
	e[1] = 'm'
	fmt.Printf("%s \n", d) // [roam] <-- Impacting the orginial array value
	fmt.Printf("%s \n", e) // [am]

	// A slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic, 
	// just as when indexing outside the bounds of a slice or array. Similarly, 
	// slices cannot be re-sliced below zero to access earlier elements in the array.

	s2 := e[3:]
	fmt.Printf("%s \n", s2) // panic: runtime error: slice bounds out of range [3:2]
}
