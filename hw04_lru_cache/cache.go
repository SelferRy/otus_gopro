package hw04lrucache

import "sync"

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
	keyMap   map[*ListItem]Key
	mu       sync.Mutex
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	_, ok := lc.items[key]
	if ok {
		lc.queue.Remove(lc.items[key])
	}
	lc.queue.PushFront(value)
	item := lc.queue.Front()
	lc.items[key] = item
	lc.keyMap[item] = key
	if lc.queue.Len() > lc.capacity {
		oldest := lc.queue.Back()
		lc.queue.Remove(oldest)
		key := lc.keyMap[oldest]
		delete(lc.items, key)
		delete(lc.keyMap, oldest)
	}
	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	lc.mu.Lock() // necessary because if ok then lc.queue will change
	defer lc.mu.Unlock()
	item, ok := lc.items[key]
	if ok {
		lc.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	lc.mu.Lock()
	defer lc.mu.Unlock()
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
		keyMap:   make(map[*ListItem]Key, capacity),
		mu:       sync.Mutex{},
	}
}
