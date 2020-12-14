package LinkedList 

import (
    "fmt"
)

type LLNode struct {
    data int 
    next *LLNode
}

type LinkedList struct {
    head *LLNode
    tail *LLNode 
    length int 
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func Test() {
    fmt.Println("Hello from LinkedList")
}

func Append(int i) {
   
}

func set(int index, int i) {

}

func (llist *LinkedList) IsEmpty() bool {
    return llist.length == 0 
}


