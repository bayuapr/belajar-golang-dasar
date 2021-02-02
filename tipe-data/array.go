package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Bayu"
	names[1] = "Apriansyah"
	names[2] = "Kaktus"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	values[0] = 100
	fmt.Println(values)
}
