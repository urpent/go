package linklist

import (
	"testing"

	"github.com/urpent/go/ut"
)

func Test_DoublyLinkedList(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		linkedList := NewDoublyLinkedList[string]()

		linkedList.Remove(nil) //do nothing

		node2 := linkedList.AddEnd("2")
		node3 := linkedList.AddEnd("3")
		ut.AssertEqual(t, "2 -> 3", linkedList.string())

		node1 := linkedList.AddFront("1")
		ut.AssertEqual(t, "1 -> 2 -> 3", linkedList.string())
		listLengthResult := linkedList.Len()
		ut.AssertEqual(t, 3, listLengthResult)

		linkedList.Remove(node2)
		ut.AssertEqual(t, "1 -> 3", linkedList.string())
		ut.AssertEqual(t, 2, linkedList.length)

		linkedList.Remove(node3)
		ut.AssertEqual(t, "1", linkedList.string())
		ut.AssertEqual(t, 1, linkedList.length)

		linkedList.Remove(node1)
		ut.AssertEqual(t, "", linkedList.string())
		ut.AssertEqual(t, 0, linkedList.length)

		ut.AssertEqual(t, (*Node[string])(nil), linkedList.head)
		ut.AssertEqual(t, (*Node[string])(nil), linkedList.tail)

		linkedList.AddFront("3")
		node2 = linkedList.AddFront("2")
		node1 = linkedList.AddFront("1")
		ut.AssertEqual(t, "1 -> 2 -> 3", linkedList.string())

		linkedList.RemoveLast()
		ut.AssertEqual(t, "1 -> 2", linkedList.string())
		linkedList.Remove(node1)
		ut.AssertEqual(t, "2", linkedList.string())
		linkedList.Remove(node2)
		ut.AssertEqual(t, "", linkedList.string())
		ut.AssertEqual(t, 0, linkedList.length)

		ut.AssertEqual(t, (*Node[string])(nil), linkedList.head)
		ut.AssertEqual(t, (*Node[string])(nil), linkedList.tail)
	})
}

func Test_DoublyLinkedList_MoveNodeToFront(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		linkedList := NewDoublyLinkedList[string]()

		linkedList.MoveNodeToFront(nil) // do nothing

		linkedList.AddEnd("2")
		linkedList.AddEnd("3")
		ut.AssertEqual(t, "2 -> 3", linkedList.string())

		node1 := linkedList.AddEnd("1")
		ut.AssertEqual(t, "2 -> 3 -> 1", linkedList.string())
		linkedList.MoveNodeToFront(node1)
		ut.AssertEqual(t, "1 -> 2 -> 3", linkedList.string())
	})
}
