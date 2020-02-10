package singlylinkedlist

import (
	"fmt"
)

//Element is a item inside a linked list.
type Element struct {
	value int
	next  *Element
}

//LinkedList holds the head element of our list.
type LinkedList struct {
	Head   *Element
	length int
	Last   *Element
}

//CreateElement creates an Element with (value)
func CreateElement(value int) Element {
	ele := Element{value, nil}
	return ele
}

//CreateList creates a new linked list.
func CreateList() LinkedList {
	ll := &LinkedList{}
	return *ll
}

//Append adds an item to the end of a linked list.
func (ll *LinkedList) Append(newEle *Element) {
	if ll.length == 0 {
		ll.Head = newEle
		ll.Last = newEle
	} else {
		lastPost := ll.Last
		lastPost.next = newEle
		ll.Last = newEle
	}
	ll.length++
}

//RemoveFirst removes the first element of our linkedlist.
func (ll *LinkedList) RemoveFirst() {
	if ll.length == 0 {

	} else if ll.length == 1 {
		ll.Head = nil
		ll.Last = nil
	} else {
		ll.Head = ll.Head.next
		ll.length--
	}
}

// RemoveLast removes the last element from a list.
func (ll *LinkedList) RemoveLast() {
	if ll.length == 0 {

	} else if ll.length == 1 {
		ll.Head = nil
		ll.Last = nil
	} else {
		last := ll.Head
		for last.next != ll.Last {
			last = last.next
		}
		last.next = nil
		ll.Last = last
		ll.length--
	}
}

// Print prints our linked list.
func (ll *LinkedList) Print() {
	if ll.length == 0 {

	} else {
		curr := ll.Head
		for curr.next != nil {
			fmt.Println(curr.value)
			curr = curr.next
		}
		fmt.Println(curr.value)
	}
}
