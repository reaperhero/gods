package main

import (
	sll "github.com/reaperhero/gods/lists/singlylinkedlist"
	"github.com/reaperhero/gods/utils"
)

// 双向链表
func main() {
	list := sll.New()
	list.Add("a")                         // ["a"]
	list.Append("b")                      // ["a","b"] (same as Add())
	list.Prepend("c")                     // ["c","a","b"]
	list.Sort(utils.StringComparator)     // ["a","b","c"]
	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Remove(2)                        // ["a","b"]
	list.Remove(1)                        // ["a"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
}
