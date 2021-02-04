package main

import "fmt"

func Ops(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ops"
	}
}

func main() {
	var data interface{} = Ops(1)
	fmt.Println(data)
}

// interface kosong digunakan ketika semua type data ingin diterima didalam parameter
// yang tidak peduli type data apapun itu
