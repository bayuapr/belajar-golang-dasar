package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Bayu Apriansyah", "Bayu"))
	fmt.Println(strings.Contains("Bayu Apriansyah", "yupp"))

	fmt.Println(strings.Split("Bayu Apriansyah", " "))
	fmt.Println(strings.ToLower("Bayu Apriansyah"))
	fmt.Println(strings.ToUpper("Bayu Apriansyah"))
	fmt.Println(strings.ToTitle("bayu apriansyah"))

	fmt.Println(strings.Trim("   Bayu  ", " "))
	fmt.Println(strings.ReplaceAll("Bay Bay Bay", "Bay", "Kaktus"))

}
