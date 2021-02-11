package main

import (
	"fmt"
	"sort"
)

type User struct {
	Nama string
	Umur int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Umur < value[j].Umur
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
func main() {
	users := []User{
		{"Bay", 22},
		{"Bayu", 25},
		{"Kaktus", 20},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
