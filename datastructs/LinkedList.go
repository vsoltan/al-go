package datastructs 

import (
    "fmt"
)

type LLNode struct {
    data int 
    *next LLNode
}

type LinkedList struct {
    LLNode *head
    LLNode *tail 
    length int 
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func Test() {
    fmt.Println("Hello from LinkedList")
}

//func Append(int i) {
//    
//}

//func set(int index, int i) {
//
//}

func IsEmpty() bool {
    return int == nil 
}


