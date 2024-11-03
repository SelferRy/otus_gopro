package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)

		l.PushFront("a")
		require.Equal(t, "a", l.Front().Value)

		l.PushBack(3.14)
		require.Equal(t, 3.14, l.Back().Value)

		latest := l.Back()
		l.MoveToFront(l.Back())
		require.Equal(t, latest, l.Front())
	})
}

func TestList_MoveToFront(t *testing.T) {
	t.Run("single value", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)
		item := l.Front()
		l.MoveToFront(item)
		require.Equal(t, item.Value, l.Front().Value)
	})

	t.Run("multiple value", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)
		l.PushFront(20)
		item := l.Back()
		l.MoveToFront(item)
		require.Equal(t, item.Value, l.Front().Value)
	})

	t.Run("first with multiple value", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)
		l.PushFront(20)
		item := l.Front()
		l.MoveToFront(item)
		require.Equal(t, item.Value, l.Front().Value)
	})
}

func TestList_Remove(t *testing.T) {
	t.Run("remove elem", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)
		first := l.Front()
		l.Remove(first)
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
		require.Equal(t, 0, l.Len())
	})
}
