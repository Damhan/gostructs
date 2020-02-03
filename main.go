package main

import (
	"fmt"
	singlylinkedlist "gostructs/Lists/SinglyLinkedList"
)

func main() {
	fmt.Println("yo")
	ll := singlylinkedlist.CreateList()
	fmt.Println(ll)
	ele := singlylinkedlist.CreateElement(5)
	ll.Append(&ele)
	fmt.Printf("List: %v\n", ll)
	fmt.Printf("Head: %v\n", ll.Head)
}
