package main

import "fmt"

func getNamaLengkap() (string, string, string) {
	return "Bayu", "Apriansyah", "kai"
}

func main() {
	namaDepan, _, namabelakang := getNamaLengkap()
	fmt.Println(namaDepan)
	fmt.Println(namabelakang)
}
