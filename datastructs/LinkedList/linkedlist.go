package LinkedList 

import (
    "strings"
    "strconv"
)

type LLNode struct {
    data int 
    next *LLNode
}

type LinkedList struct {
    front *LLNode
    back *LLNode 
    length int 
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func (list *LinkedList) IsEmpty() bool {
    return list.length == 0 
}

func (list *LinkedList) AppendFront(i int) {
    to_append := &LLNode{i, nil}
    if list.front == nil {
        list.front = to_append
        list.back = to_append 
    } else {
        to_append.next = list.front 
        list.front = to_append
    }
    list.length++ 
}

func (list *LinkedList) AppendBack(i int) {
    to_append := &LLNode{i, nil}
    if list.back == nil {
        list.front = to_append
        list.back = to_append 
    } else {
        list.back.next = to_append  
        list.back = to_append 
    }
    list.length++ 
}

func (list *LinkedList) RemoveFront() int {
    if list.length == 0 {
        // error 
    }
    removed_data := list.front.data 
    list.front = list.front.next 
    list.length-- 
    return removed_data
}

func (list *LinkedList) RemoveBack() int {
    if list.length == 0 {
        // error 
    }
    removed_data := list.back.data 
    it := list.front
    for it.next != list.back {
        it = it.next 
    }
    list.back = it 
    list.back.next = nil 
    list.length-- 
    return removed_data
}

func (list *LinkedList) String() string {
    it := list.front; 
    var list_str strings.Builder 
    for it != nil {
        list_str.WriteString(strconv.Itoa(it.data))
        list_str.WriteString(" ")
        it = it.next 
    }
    return list_str.String()
}
