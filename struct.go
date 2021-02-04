package main

import "fmt"

type Pelanggan struct {
	nama, alamat string
	umur         int
}

func (pelanggan Pelanggan) cok(nama string) {
	fmt.Println("Woi", nama, "kenalin nama saya", pelanggan.nama)
}

func (x Pelanggan) cuy() {
	fmt.Println("Cuuy ", x.nama)
}

func main() {
	//lebih disarankan
	var bay Pelanggan
	bay.nama = "bayu"
	bay.alamat = "Mars"
	bay.umur = 22

	bay.cok("Yub")
	bay.cuy()
	// fmt.Println(bay)
	// fmt.Println(bay.nama)
	// fmt.Println(bay.alamat)
	// fmt.Println(bay.umur)

	// //lebih disarankan
	// kaktus := Pelanggan{
	// 	nama:   "Kaktus",
	// 	alamat: "Tandus",
	// 	umur:   100,
	// }
	// fmt.Println(kaktus)

	// moskov := Pelanggan{"Mooskov", "Pluto", 2020}
	// fmt.Println(moskov)
}
