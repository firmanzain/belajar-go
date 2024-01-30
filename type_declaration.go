package main

import "fmt"

func main() {
	type NoKTP string

	var ktp NoKTP = "111111"
	var ktp2 string = "222222"
	var ktp3 NoKTP = NoKTP(ktp2)

	fmt.Println(ktp)
	fmt.Println(ktp3)
}
