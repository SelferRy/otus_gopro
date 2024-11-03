package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := ListItem{v, l.head, nil}
	if l.head == nil {
		l.head = &item
		l.tail = &item
	} else {
		l.head.Prev = &item
		l.head = &item
	}
	l.len++
	return &item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := ListItem{v, nil, l.tail}
	if l.head == nil {
		l.head = &item
		l.tail = &item
	} else {
		l.tail.Next = &item
		l.tail = &item
	}
	l.len++
	return &item
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
		l.head.Prev = nil
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
		l.tail.Next = nil
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	i.Prev = nil
	i.Next = l.head
	l.head.Prev = i
	l.head = i
}

func NewList() List {
	return new(list)
}
