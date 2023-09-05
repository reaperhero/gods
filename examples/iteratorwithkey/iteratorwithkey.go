package main

import (
	"fmt"
	"github.com/reaperhero/gods/maps/treemap"
)

// set with key
// Next() get value by sort by desc
// Prev() get value by sort by acs
func main() {
	m := treemap.NewWithIntComparator()
	m.Put(1, "a")
	m.Put(2, "b")
	m.Put(3, "b")
	it := m.Iterator()

	fmt.Print("\nForward iteration\n")
	for it.Next() {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]") // [0:a][1:b][2:c]
	}

	fmt.Print("\nBackward iteration\n")
	for it.Prev() {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]") // [2:c][1:b][0:a]
	}
}
