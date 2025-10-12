package event

import "sync"

type EventHandler func(data interface{})

type EventBus struct {
	subscribers map[string][]EventHandler
	lock        sync.Mutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]EventHandler),
	}
}

func (b *EventBus) Subscribe(event string, handler EventHandler) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.subscribers[event] = append(b.subscribers[event], handler)
}

func (b *EventBus) Publish(event string, data interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if handlers, ok := b.subscribers[event]; ok {
		for _, handler := range handlers {
			go handler(data)
		}
	}
}
