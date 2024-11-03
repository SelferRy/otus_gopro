package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	_, ok := lc.items[key]
	if ok {
		lc.queue.Remove(lc.items[key])
	}
	lc.queue.PushFront(value)
	lc.items[key] = lc.queue.Front()
	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := lc.items[key]
	if ok {
		lc.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	lc.queue = NewList()
	threshold := 10000 // the constant can be changed
	if len(lc.items) < threshold {
		clear(lc.items)
	} else {
		lc.items = make(map[Key]*ListItem, lc.capacity)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
