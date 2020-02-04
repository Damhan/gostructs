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
	ele2 := singlylinkedlist.CreateElement(10)
	ele3 := singlylinkedlist.CreateElement(6)
	ll.Append(&ele)
	ll.Append(&ele2)
	ll.Append(&ele3)
	fmt.Printf("List: %v\n", ll)
	fmt.Printf("Head: %v\n", ll.Head)
	fmt.Printf("Last: %v\n", ll.Last)
	ll.RemoveLast()
	fmt.Printf("After Removal: %v\n", ll)
}
