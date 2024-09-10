package parcial1MAJU2MA

//Implementar la primitiva (método) de DoubleLinkedList, func (l *DoubleLinkedList[T])Extend(other *DoubleLinkedList[T]){..} que extienda la lista con todos los elementos de la otra lista que se pasa por parámetro. Indicar y justificar el orden del algoritmo. Al finalizar la ejecución la lista sobre la que se ejecutó Extend debe tener al final de la misma concatenada la lista other y la lista other debe quedar vacía. El método debe ser O(1)


import(
	"fmt"
	l "github.com/trigologiaa/data-structures/list"
)

type DoubleLinkedList[T comparable] struct {
	head *l.DoubleLinkedNode[T]
	tail *l.DoubleLinkedNode[T]
	size int
}

func NewDoubleLinkedList[T comparable]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

func (li *DoubleLinkedList[T]) Head() *l.DoubleLinkedNode[T] {
	return li.head
}

func (li *DoubleLinkedList[T]) Tail() *l.DoubleLinkedNode[T] {
	return li.tail
}

func (li *DoubleLinkedList[T]) Size() int {
	return li.size
}

func (li *DoubleLinkedList[T]) IsEmpty() bool {
	return li.size == 0
}

func (li *DoubleLinkedList[T]) Clear() {
	li.head = nil
	li.tail = nil
	li.size = 0
}

func (li *DoubleLinkedList[T]) Prepend(data T) {
	newNode := l.NewDoubleLinkedNode[T](data)
	if li.size == 0 {
		li.tail = newNode
	} else {
		newNode.SetNext(li.head)
		li.head.SetPrev(newNode)
	}
	li.head = newNode
	li.size++
}

func (li *DoubleLinkedList[T]) Append(data T) {
	newNode := l.NewDoubleLinkedNode[T](data)
	if li.size == 0 {
		li.head = newNode
	} else {
		li.tail.SetNext(newNode)
		newNode.SetPrev(li.tail)
	}
	li.tail = newNode
	li.size++
}

func (li *DoubleLinkedList[T]) Find(data T) *l.DoubleLinkedNode[T] {
	for current := li.head; current != nil; current = current.Next() {
		if current.Data() == data {
			return current
		}
	}

	return nil
}

func (li *DoubleLinkedList[T]) RemoveFirst() {
	if li.IsEmpty() {
		return
	}

	li.head = li.head.Next()
	li.size--

	if li.IsEmpty() {
		li.tail = nil
	} else {
		li.head.SetPrev(nil)
	}
}

func (li *DoubleLinkedList[T]) RemoveLast() {
	if li.IsEmpty() {
		return
	}

	if li.size == 1 {
		li.head = nil
		li.tail = nil
		li.size = 0

		return
	}

	li.tail = li.tail.Prev()
	li.tail.SetNext(nil)
	li.size--
}

func (li *DoubleLinkedList[T]) Remove(data T) {
	node := li.Find(data)

	if node == nil {
		return
	}

	if node == li.head {
		li.RemoveFirst()

		return
	}

	if node == li.tail {
		li.RemoveLast()

		return
	}

	node.Prev().SetNext(node.Next())
	node.Next().SetPrev(node.Prev())
	li.size--
}

func (li *DoubleLinkedList[T]) String() string {
	if li.IsEmpty() {
		return "DoubleLinkedList: []"
	}

	result := "DoubleLinkedList: "

	current := li.Head()
	for {
		result += fmt.Sprintf("[%v]", current.Data())
		if !current.HasNext() {
			break
		}
		result += " ↔ "
		current = current.Next()
	}

	return result
}

func(li *DoubleLinkedList[T]) Extend(other *l.DoubleLinkedList[T]) {
	if other.IsEmpty() {
		return
	}
	if li.IsEmpty() {
		li.head = other.Head()
		li.tail = other.Tail()
		li.size = other.Size()
	} else {
		li.tail.SetNext(other.Head())
		other.Head().SetPrev(li.tail)
		li.tail = other.Tail()
		li.size += other.Size()
	}
	other.Clear()
}