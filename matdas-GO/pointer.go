package main

import "fmt"

type Alamat struct {
	Kota, Provinsi, Negara string
}

func UbahNegaraKeIndonesia(alamat *Alamat) {
	alamat.Negara = "Indonesia"
}

func main() {
	//pas by value
	var alamat1 Alamat = Alamat{"Yogyakarta", "D.I. Yogyakarta", "Indonesia"}
	//pas by reference menggunkan &
	//pas by value tanpa &
	var alamat2 *Alamat = &alamat1
	var alamat3 *Alamat = &alamat1

	alamat2.Kota = "Sleman"

	*alamat2 = Alamat{"Makassar", "SULSEL", "Indonesia"}

	fmt.Println(alamat1) //alamat tidak berubah
	fmt.Println(alamat2)
	fmt.Println(alamat3)

	var alamat4 *Alamat = new(Alamat)
	alamat4.Kota = "Bulukumba"
	fmt.Println(alamat4)

	var alamat = Alamat{
		Kota:     "Jakarta",
		Provinsi: "DKI Jakarta",
		Negara:   "",
	}
	var alamatPointer *Alamat = &alamat
	UbahNegaraKeIndonesia(alamatPointer)
	fmt.Println(alamat)

}
