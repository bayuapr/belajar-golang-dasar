package main

import "fmt"

func main() {
	nama := "Bayu"
	counter := 0

	increment := func() {
		nama := "Kaktus"
		fmt.Println("Increment")
		counter++
		fmt.Println(nama)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(nama)
}

//closure mengambil data dari atas kebawah
