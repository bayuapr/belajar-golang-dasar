package main

import "fmt"

func main() {
	//variable nama
	var nama = "Bayu"

	if nama == "Bayu" {
		//jika benar
		fmt.Println("Hello Bay")
	} else if nama == "Kaktus" {
		//jika kondisi ke-2
		fmt.Println("Kaktusnesia")
	} else {
		//jika salah
		fmt.Println("Maaf, Anda Salah Input")
	}

	if length := len(nama); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah betul")
	}
}
