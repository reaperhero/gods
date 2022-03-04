package main

import "github.com/reaperhero/gods/sets/linkedhashset"

// set by item
// 根据插入顺序排序
func main() {
	set := linkedhashset.New() // empty
	set.Add(5)                 // 5
	set.Add(4, 4, 3, 2, 1)     // 5, 4, 3, 2, 1 (in insertion-order, duplicates ignored)
	set.Remove(4)              // 5, 3, 2, 1 (in insertion-order)
	set.Remove(2, 3)           // 5, 1 (in insertion-order)
	set.Contains(1)            // true
	set.Contains(1, 5)         // true
	set.Contains(1, 6)         // false
	_ = set.Values()           // []int{5, 1} (in insertion-order)
	set.Clear()                // empty
	set.Empty()                // true
	set.Size()                 // 0
}
