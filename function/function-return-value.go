package main

import "fmt"

func getHello(nama string) string {
	if nama == "" {
		return "Hello Ngab"
	} else {
		return "Hello " + nama
	}
}

func main() {
	result := getHello("Bay")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
