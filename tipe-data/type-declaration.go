package main

import "fmt"

func main() {
	type NoKTM string
	type Menikah bool

	var noKtmKaktus NoKTM = "1234567890987654"
	var menikahStatus Menikah = true
	fmt.Println(noKtmKaktus)
	fmt.Println(menikahStatus)

}
