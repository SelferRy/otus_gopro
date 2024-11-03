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
	keyMap   map[*ListItem]Key
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
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
		keyMap:   make(map[*ListItem]Key, capacity),
	}
}
