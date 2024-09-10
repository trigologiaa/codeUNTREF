package parcial1MAJU2TA

//Implementar la primitiva (método) de CircularList: func (l *CircularList[T])InsertAt(data T, pos int){..} que inserte el dato dado en la posición pos % size, dónde size es el tamaño actual de la lista. Indicar el orden del método. La lista está implementada sobre nodos doblemente enlazados.


import(
	"fmt"
	l "github.com/trigologiaa/data-structures/list"
)

type CircularList[T comparable] struct {
	head *l.DoubleLinkedNode[T]
	size int
}

func NewCircularList[T comparable]() *CircularList[T] {
	return &CircularList[T]{}
}

func (li *CircularList[T]) Head() *l.DoubleLinkedNode[T] {
	return li.head
}

func (li *CircularList[T]) Tail() *l.DoubleLinkedNode[T] {
	if li.size == 0 {
		return nil
	}

	return li.head.Prev()
}

func (li *CircularList[T]) Size() int {
	return li.size
}

func (li *CircularList[T]) IsEmpty() bool {
	return li.size == 0
}

func (li *CircularList[T]) Clear() {
	li.head = nil
	li.size = 0
}

func (li *CircularList[T]) Prepend(data T) {
	node := l.NewDoubleLinkedNode(data)
	if li.size == 0 {
		li.head = node
		li.head.SetNext(li.head)
		li.head.SetPrev(li.head)
	} else {
		node.SetNext(li.head)
		node.SetPrev(li.head.Prev())
		li.head.Prev().SetNext(node)
		li.head.SetPrev(node)
		li.head = node
	}
	li.size++
}

func (li *CircularList[T]) Append(data T) {
	node := l.NewDoubleLinkedNode(data)
	if li.size == 0 {
		li.head = node
		li.head.SetNext(li.head)
		li.head.SetPrev(li.head)
	} else {
		node.SetNext(li.head)
		node.SetPrev(li.head.Prev())
		li.head.Prev().SetNext(node)
		li.head.SetPrev(node)
	}
	li.size++
}

func (li *CircularList[T]) Find(data T) *l.DoubleLinkedNode[T] {
	if li.size == 0 {
		return nil
	}

	node := li.head
	for i := 0; i < li.size; i++ {
		if node.Data() == data {
			return node
		}
		node = node.Next()
	}

	return nil
}

func (li *CircularList[T]) RemoveFirst() {
	if li.size == 0 {
		return
	}

	if li.size == 1 {
		li.head = nil
		li.size--

		return
	}

	li.head.Prev().SetNext(li.head.Next())
	li.head.Next().SetPrev(li.head.Prev())
	li.head = li.head.Next()
	li.size--
}

func (li *CircularList[T]) RemoveLast() {
	if li.size == 0 {
		return
	}

	if li.size == 1 {
		li.head = nil
		li.size--

		return
	}

	li.head.Prev().Prev().SetNext(li.head)
	li.head.SetPrev(li.head.Prev().Prev())
	li.size--
}

func (li *CircularList[T]) Remove(data T) {
	node := li.Find(data)
	if node == nil {
		return
	}

	if node == li.head {
		li.RemoveFirst()

		return
	}

	node.Prev().SetNext(node.Next())
	node.Next().SetPrev(node.Prev())
	li.size--
}

func (li *CircularList[T]) String() string {
	if li.IsEmpty() {
		return "CircularList: ⇢ [] ⇠"
	}

	result := "CircularList: ⇢ "

	current := li.Head()
	for {
		result += fmt.Sprintf("[%v]", current.Data())
		if current == li.Tail() {
			break
		}
		result += " ↔ "
		current = current.Next()
	}
	result += " ⇠"

	return result
}

func (li *CircularList[T]) InsertAt(data T, pos int) {
	if pos < 0 {
		pos = li.size + (pos % li.size)
	} else {
		pos = pos % li.size
	}
	nuevoNodo := l.NewDoubleLinkedNode(data)
	if li.size == 0 {
		li.head = nuevoNodo
		li.head.SetNext(li.head)
		li.head.SetPrev(li.head)
	} else {
		actual := li.head
		for i := 0; i < pos; i++ {
			actual = actual.Next()
		}
		nuevoNodo.SetNext(actual)
		nuevoNodo.SetPrev(actual.Prev())
		actual.Prev().SetNext(nuevoNodo)
		actual.SetPrev(nuevoNodo)
		if pos == 0 {
			li.head = nuevoNodo
		}
	}
	li.size++
}