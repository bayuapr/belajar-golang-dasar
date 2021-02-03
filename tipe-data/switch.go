package main

import "fmt"

func main() {
	// kondisi data
	nama := "Kaktus"

	switch nama {
	// jika ada
	case "Bayu":
		fmt.Println("Hello Bay")
		fmt.Println("Hello Bay")
	// jika ada
	case "Kaktus":
		fmt.Println("Hello Kaktus")
		fmt.Println("Hello Kaktus")
	// jika tidak ada
	default:
		fmt.Println("Kenalan kuy")
		fmt.Println("Kenalan kuy")
	}

	// mengecek panjang ukuruan dan panjang tidak boleh lebih dari 5
	// switch length := len(nama); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")

	// }

	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama terlalu panjang")
	}
}
