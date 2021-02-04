package main

import "fmt"

func random() interface{} {
	return "kosong..."
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "ini string")
	case int:
		fmt.Println("Value", value, "ini int")
	default:
		fmt.Println("Tipe tidak diketahui")
	}
}
