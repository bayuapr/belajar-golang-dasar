package main

import "fmt"

func main() {

	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// bulan[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei Lagi"
	// fmt.Println(bulan)

	var slice2 = bulan[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Bayu")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(bulan)
}
