package main

import "fmt"

func endApp() {
	massage := recover()
	if massage != nil {
		fmt.Println("Error dengan masssage:", massage)

	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi sedang berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Bay")
}
