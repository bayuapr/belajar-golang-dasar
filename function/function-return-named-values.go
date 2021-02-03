package main

import "fmt"

func getNamaLengkap2() (namadepan string, namatengah string, namabelakang string) {
	namadepan = "Mas"
	namatengah = "Bayu"
	namabelakang = "Apriansyah"

	return
}

func main() {
	x, y, z := getNamaLengkap2()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
