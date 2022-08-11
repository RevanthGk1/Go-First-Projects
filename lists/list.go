package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	// store the reference
	v := l.PushBack(10)
	fmt.Println(l.Front()) // &{0x43e280 0x43e280 0x43e280 10}

	l.PushBack(12)
	l.PushBack(14)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}

	// remove using the reference
	l.Remove(v)
	fmt.Println(l.Front()) // <nil>
}
