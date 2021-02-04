package main

import "fmt"

type HasNama interface {
	GetNama() string
}

func Haii(hasNama HasNama) {
	fmt.Println("Haaii", hasNama.GetNama())
}

type Person struct {
	Nama string
}

func (person Person) GetNama() string {
	return person.Nama
}

type Hewan struct {
	Nama string
}

func (hewan Hewan) GetNama() string {
	return hewan.Nama
}

func main() {
	var bayu Person
	bayu.Nama = "Bay"

	Haii(bayu)

	cat := Hewan{
		Nama: "Miong",
	}
	Haii(cat)
}
