package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Ganti host")
	var user *string = flag.String("user", "root", "Ganti database user")
	var password *string = flag.String("password", "root", "Ganti database password")

	flag.Parse()
	fmt.Println("Host : ", host)
	fmt.Println("User : ", user)
	fmt.Println("Password : ", password)
}
