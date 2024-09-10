package recuParcial1MAMI

//Agregar a la implementación de lista enlazada simple un método denominado Revertir, que revierta los elementos de la lista, tal que el primer nodo pase a ser el último y viceversa. Justificar el orden de la solución. func (l *LinkedList[T]) Revertir()

import(
	"fmt"
	l "github.com/trigologiaa/data-structures/list"
)

type LinkedList[T comparable] struct {
	head *l.LinkedNode[T]
	tail *l.LinkedNode[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Head() *l.LinkedNode[T] {
	return l.head
}

func (l *LinkedList[T]) Tail() *l.LinkedNode[T] {
	return l.tail
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (li *LinkedList[T]) Prepend(data T) {
	newNode := l.NewLinkedListNode(data)
	if li.IsEmpty() {
		li.tail = newNode
	} else {
		newNode.SetNext(li.head)
	}
	li.head = newNode
	li.size++
}

func (li *LinkedList[T]) Append(data T) {
	newNode := l.NewLinkedListNode(data)
	if li.IsEmpty() {
		li.head = newNode
	} else {
		li.tail.SetNext(newNode)
	}
	li.tail = newNode
	li.size++
}

func (l *LinkedList[T]) Find(data T) *l.LinkedNode[T] {
	for current := l.head; current != nil; current = current.Next() {
		if current.Data() == data {
			return current
		}
	}

	return nil
}

func (l *LinkedList[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}

	l.head = l.head.Next()

	if l.head == nil {
		l.tail = nil
	}

	l.size--
}

func (l *LinkedList[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}

	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		current := l.head
		for current.Next() != l.tail {
			current = current.Next()
		}
		current.SetNext(nil)
		l.tail = current
	}
	l.size--
}

func (l *LinkedList[T]) Remove(data T) {
	node := l.Find(data)

	if node == nil {
		return
	}

	if node == l.head {
		l.RemoveFirst()

		return
	}

	current := l.Head()

	for current.Next() != node {
		current = current.Next()
	}

	current.SetNext(node.Next())

	if node == l.tail {
		l.tail = current
	}
	l.size--
}

func (l *LinkedList[T]) String() string {
	if l.IsEmpty() {
		return "LinkedList: []"
	}

	result := "LinkedList: "

	current := l.Head()
	for {
		result += fmt.Sprintf("[%v]", current.Data())
		if !current.HasNext() {
			break
		}
		result += " → "
		current = current.Next()
	}

	return result
}

func(li *LinkedList[T]) Revertir() {
	if li.IsEmpty() {
		return
	}
	var previo, siguiente *l.LinkedNode[T]
	actual := li.head
	li.tail = actual
	for actual != nil {
		siguiente = actual.Next()
		actual.SetNext(previo)
		previo = actual
		actual = siguiente
	}
	li.head = previo
}