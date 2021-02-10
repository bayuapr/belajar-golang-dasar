package main

import "fmt"

type Blacklist func(string) bool

func registeruser(nama string, blacklist Blacklist) {
	if blacklist(nama) {
		fmt.Println("Kamu telah diblokir", nama)
	} else {
		fmt.Println("Selamat", nama)
	}
}

// func blacklistAdmin(nama string) bool {
// 	return nama == "admin"
// }

// func blacklistRoot(nama string) bool {
// 	return nama == "root"
// }
func main() {
	blacklist := func(nama string) bool {
		return nama == "admin"
	}

	registeruser("admin", blacklist)
	registeruser("bay", blacklist)

	registeruser("root", func(nama string) bool {
		return nama == "root"
	})
	registeruser("bay", func(nama string) bool {
		return nama == "root"
	})
}
