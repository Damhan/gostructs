package singlylinkedlist

//Element is a item inside a linked list.
type Element struct {
	value int
	next  *Element
}

//LinkedList holds the head element of our list.
type LinkedList struct {
	Head   *Element
	length int
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
	} else {
		currentEle := ll.Head
		for currentEle.next != nil {
			currentEle = currentEle.next
		}
		currentEle.next = newEle
	}
	ll.length++
}
