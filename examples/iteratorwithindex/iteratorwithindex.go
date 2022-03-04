package main

import (
	"fmt"
	"github.com/reaperhero/gods/sets/treeset"
)

// set with key
// Next() get value by sort by desc
// Prev() get value by sort by acs
func main() {
	set := treeset.NewWithStringComparator()
	set.Add("a", "b", "c")
	it := set.Iterator()

	fmt.Print("\nForward iteration\n")
	for it.Next() {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]") // [0:a][1:b][2:c]
	}

	fmt.Print("\nBackward iteration\n")
	for it.Prev() {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]") // [2:c][1:b][0:a]
	}

	fmt.Print("\nBackward iteration (again)\n")
	for it.End(); it.Prev(); {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]") // [2:c][1:b][0:a]
	}

}
