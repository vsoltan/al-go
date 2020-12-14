package Queue 

import (
	"github.com/vsoltan/al-go/datastructs/LinkedList"
)

type Queue struct {
	list *LinkedList.LinkedList 
}

func New() *Queue {
	return &Queue{list: LinkedList.New()}
}

func (q *Queue) Enqueue(value interface{}) {
	q.list.AppendBack(value)
}

func (q *Queue) Dequeue() (value interface{}, ok bool) {
	if (q.Size() <= 0) {
		return 
	}
	value, ok = q.list.RemoveFront()
	return 
}

func (q *Queue) Peek() (value interface{}, ok bool) {
	if (q.Size() <= 0) {
		return 
	}
	value = q.list.Front().Value() 
	ok = true 
	return 
}

func (q *Queue) Size() int {
	return q.list.Length() 
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}