package parcial1MIVI2

//Implementar una primitiva (método) de LinkedList, Interleave[T comparable](other *LinkedList[T]), que intercale los nodos de la lista (receptor) y other en la lista actual. Indicar y justificar el orden del algoritmo. Al finalizar la ejecución la lista sobre la que se ejecutó Interaleave debe tener el tamaño de la lista original más el tamaño de other y la lista other debe quedar vacía.

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

func(l *LinkedList[T]) Interleave(other *LinkedList[T]) {
	if other.IsEmpty() {
		return
	}
	actual := l.Head()
	otherActual := other.Head()
	for actual != nil && otherActual != nil {
		siguienteActual := actual.Next()
		siguienteOtherActual := otherActual.Next()
		actual.SetNext(otherActual)
		otherActual.SetNext(siguienteActual)
		actual = siguienteActual
		otherActual = siguienteOtherActual
	}
	if otherActual != nil {
		l.tail.SetNext(otherActual)
		l.tail = other.tail
	}
	l.size += other.Size()
	other.head = nil
	other.tail = nil
	other.size = 0
}

/*
Justificación del Orden del Algoritmo
El algoritmo implementado recorre ambas listas una sola vez, por lo tanto, su complejidad temporal es O(n + m), donde n es el tamaño de la lista l y m es el tamaño de la lista other. Esto se debe a que estamos realizando un único recorrido simultáneo de ambas listas, intercalando nodos de manera directa y efectuando ajustes finales en caso de que other tenga más nodos que l.
Consideraciones Finales
Este método modifica la lista actual (l) y vacía la lista other al finalizar.
Se asume que other es del mismo tipo (LinkedList[T]) para facilitar el acceso a sus métodos y propiedades.
Es importante asegurar que la operación no cause ciclos o inconsistencias en la estructura de datos enlazada.
Con esta implementación, puedes intercalar los nodos de dos listas enlazadas, manteniendo la eficiencia y asegurando que ambas listas queden en un estado coherente después de la operación.
*/