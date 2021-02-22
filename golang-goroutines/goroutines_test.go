package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHaloKaktus() {
	fmt.Println("Halo Kaktus")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHaloKaktus()
	fmt.Println("Hmmm")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
