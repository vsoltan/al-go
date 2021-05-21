package linkedlist

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

func (node *Node) Value() interface{} {
	return node.value
}

func (node *Node) SetValue(value interface{}) {
	node.value = value
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) SetNext(next *Node) {
	node.next = next
}

type List struct {
	front *Node
	back  *Node
	len   int
}

func New() *List {
	return &List{}
}

func (list *List) Front() *Node {
	return list.front
}

func (list *List) Back() *Node {
	return list.back
}

func (list *List) Length() int {
	return list.len
}

func (list *List) IsEmpty() bool {
	return list.len == 0
}

func (list *List) AppendFront(i interface{}) {
	toAppend := &Node{i, nil}
	if list.front == nil {
		list.front = toAppend
		list.back = toAppend
	} else {
		toAppend.next = list.front
		list.front = toAppend
	}
	list.len++
}

func (list *List) AppendBack(i interface{}) {
	toAppend := &Node{i, nil}
	if list.back == nil {
		list.front = toAppend
		list.back = toAppend
	} else {
		list.back.next = toAppend
		list.back = toAppend
	}
	list.len++
}

func (list *List) RemoveFront() (value interface{}, ok bool) {
	if list.len == 0 {
		return
	}
	value = list.front.value
	list.front = list.front.next
	list.len--
	ok = true
	return
}

func (list *List) RemoveBack() (value interface{}, ok bool) {
	if list.len == 0 {
		return
	}
	value = list.back.value
	it := list.front
	for it.Next() != list.back {
		it = it.Next()
	}
	list.back = it
	list.back.next = nil
	list.len--
	ok = true
	return
}

func (list *List) Print() {
	it := list.front
	for it != nil {
		fmt.Println(it.value)
		it = it.Next()
	}
}

func (list *List) Reverse() *List {
	return nil
}
