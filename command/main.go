package main

import (
	"fmt"

	"github.com/clayessex/godev/vessels"
)

func main() {
	d := vessels.NewDeque[int]()
	d.PushBack(42)
	fmt.Println("Pop: ", d.PopBack())
}
