package main

import (
	"fmt"

	"github.com/clayessex/godev/vessels"
)

func main() {
	fmt.Println("testing a LRUCache")
	d := vessels.NewDeque[int]()
	d.PushBack(9)
	fmt.Println("Test pop: ", d.PopBack())
}
