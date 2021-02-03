package main

import "fmt"

type Filter func(string) string

func HelloDenganFilter(nama string, filter Filter) {
	namaFiltered := filter(nama)
	fmt.Println("Hello", namaFiltered)
}

func spamFilter(nama string) string {
	if nama == "Anying" {
		return "..."
	} else {
		return nama
	}
}

func main() {
	HelloDenganFilter("Bay", spamFilter)
	HelloDenganFilter("Anying", spamFilter)
}
