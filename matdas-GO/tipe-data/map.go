package main

import "fmt"

func main() {

	person := map[string]string{
		"nama":   "Bayu",
		"alamat": "DIY",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Go-leng"
	book["author"] = "Bayby"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
