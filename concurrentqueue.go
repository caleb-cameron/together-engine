package engine

import (
	"io"
	"sync"
)

type ConcurrentQueue struct {
	sync.RWMutex
	items []ConcurrentQueueItem
}

type ConcurrentQueueItem struct {
	Value interface{}
}

func NewConcurrentQueue() *ConcurrentQueue {
	return &ConcurrentQueue{
		items: []ConcurrentQueueItem{},
	}
}

func (q *ConcurrentQueue) Push(item interface{}) {
	q.Lock()
	defer q.Unlock()

	i := ConcurrentQueueItem{
		Value: item,
	}

	q.items = append(q.items, i)
}

func (q *ConcurrentQueue) Delete(index int) {
	q.Lock()
	defer q.Unlock()

	q.items = append(q.items[:index], q.items[index+1:]...)
}

func (q *ConcurrentQueue) Pop() (ConcurrentQueueItem, error) {
	if len(q.items) == 0 {
		return ConcurrentQueueItem{}, io.EOF
	}

	q.Lock()
	defer q.Delete(0)
	defer q.Unlock()

	return q.items[0], nil
}

func (q *ConcurrentQueue) Iter() <-chan ConcurrentQueueItem {
	c := make(chan ConcurrentQueueItem)

	f := func() {
		for {
			item, err := q.Pop()
			if err == io.EOF {
				break
			}
			c <- item
		}
		close(c)
	}
	go f()

	return c
}
