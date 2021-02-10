package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Nama string
}

func main() {
	sample := Sample{"Bay"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Nama)
}
