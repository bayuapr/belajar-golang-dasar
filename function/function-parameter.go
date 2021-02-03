package main

import "fmt"

func HelloTo(namaDepan string, namaBelakang string) {
	fmt.Println("Hello", namaDepan, namaBelakang)
}

func main() {
	namaDepan := "Bay"
	HelloTo(namaDepan, "Apriansyah")
	HelloTo("Duri", "Kaktus")
}
