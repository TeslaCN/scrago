package task

import (
	"errors"
	"net/url"
)

type Pool interface {
	Fetch() (*Task, error)
	Offer(Task)
}

type Factory interface {
	Process(Raw)
}

type Task struct {
	Method string
	Url    url.URL
}

type Raw struct {
	Url  url.URL
	Data []byte
}

type DefaultPool struct {
	queue Queue
}

func (p *DefaultPool) Offer(task Task) {
	// log.Printf("Pool <- [%s]\n", task.Url.String())
	p.queue.Put(task)
}

func (p *DefaultPool) Fetch() (*Task, error) {
	task, e := p.queue.Take()
	if e != nil {
		return nil, e
	}
	// log.Printf("Pool -> [%s]\n", task.Url.String())
	return task, nil
}

func NewPool() Pool {
	return &DefaultPool{queue: NewChannelQueue()}
}

type DefaultFactory struct {
}

func (f *DefaultFactory) Process(Raw) {

}

type Queue interface {
	Put(Task)
	Take() (*Task, error)
	Size() int
}

type ChannelQueue struct {
	channel chan Task
}

func NewChannelQueue() Queue {
	return &ChannelQueue{channel: make(chan Task, 1<<10)}
}

func (q *ChannelQueue) Size() int {
	return len(q.channel)
}

func (q *ChannelQueue) Put(task Task) {
	q.channel <- task
}

func (q *ChannelQueue) Take() (*Task, error) {
	task, ok := <-q.channel
	if !ok {
		return nil, errors.New("Fetch from channel failed")
	}
	return &task, nil
}
