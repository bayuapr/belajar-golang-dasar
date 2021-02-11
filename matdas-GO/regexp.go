package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("b([a-z])y")

	fmt.Println(regex.MatchString("bay"))
	fmt.Println(regex.MatchString("bAy"))

	search := regex.FindAllString("bayu bay uyab baybay", -1)
	fmt.Println(search)
}
