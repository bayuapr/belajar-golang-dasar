package main

import "fmt"

type Pria struct {
	Nama string
}

func (pria *Pria) Menikah() {
	pria.Nama = "Mas. " + pria.Nama
}

func main() {
	bay := Pria{"Bay"}
	bay.Menikah()

	fmt.Println(bay.Nama)
}
