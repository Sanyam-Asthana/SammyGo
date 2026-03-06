package modules

import (
	"fmt"
	"slices"
)

type Queue struct {
	Data []string
}

func NewQueue() *Queue {
	q := Queue{}
	return &q
}

func (queueRef *Queue) Enqueue(data string) {
	queueRef.Data = append(queueRef.Data, data)
}

func (queueRef *Queue) Dequeue() string {
	dequeuedElement := queueRef.Data[0]
	queueRef.Data = slices.Delete(queueRef.Data, 0, 1)
	return dequeuedElement
}

func (queueRef *Queue) Print() {
	fmt.Println(queueRef.Data)
}


