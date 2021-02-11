package main

import "fmt"

func NewMap(nama string) map[string]string {
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"nama": nama,
		}
	}
}

func main() {
	// person := NewMap("Bay")

	var person map[string]string = NewMap("Cok")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)

	}

}

//nil digunakan data kosong
