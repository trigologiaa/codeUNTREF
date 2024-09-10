package parcial1MIVI

import(
	e "errors"
)

type nodo[T any] struct {
	value T
	next *nodo[T]
}

type LinkedList[T comparable] struct {
	head *nodo[T] // puntero al primer nodo
	tail *nodo[T] // puntero al último nodo
	size int
}

func (l *LinkedList[T]) verElPrimerElemento() T {
	if l.head == nil {
		var zero T
		return zero
	}
	return l.head.value
}

func (l *LinkedList[T]) InsertarPrimero(valor T) {
	newNode := &nodo[T]{value: valor}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.size++
}

func (l *LinkedList[T]) InsertarUltimo(valor T) {
	newNode := &nodo[T]{value: valor}
	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *LinkedList[T]) RemoverPrimero() T {
	if l.size == 0 {
		var zero T
		return zero
	}
	removedValue := l.head.value
	l.head = l.head.next
	l.size--
	if l.size == 0 {
		l.tail = nil
	}
	return removedValue
}

func (l *LinkedList[T]) RemoverUltimo() T {
	if l.size == 0 {
		var zero T
		return zero
	}
	if l.size == 1 {
		removedValue := l.head.value
		l.head = nil
		l.tail = nil
		l.size = 0
		return removedValue
	}
	current := l.head
	for current.next != l.tail {
		current = current.next
	}
	removedValue := l.tail.value
	current.next = nil
	l.tail = current
	l.size--
	return removedValue
}

func (l *LinkedList[T]) Tamanio() int {
	return l.size
}

func (l *LinkedList[T]) EstaVacia() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Contiene(valor T) bool {
	current := l.head
	for current != nil {
		if current.value == valor {
			return true
		}
		current = current.next
	}
	return false
}

//Implementar una Cola (LIFO) sobre una lista enlazada simple. La cola debe soportar las siguientes operaciones: Encolar, Desencolar, Primero (para ver el que está primero en la cola) y EstaVacia. Debe soportar cualquier tipo de dato comparable y lanzar los errores que correspondan cuando se intenta desencolar o consultar el primero de una cola vacía. Todas las operaciones deben ser O(1) NOTA: No deben implementar los métodos de LinkedList, sólo deben usarlos.

type Cola[T comparable] struct {
	list *LinkedList[T]
}

func NuevaCola[T comparable]() *Cola[T] {
	return &Cola[T] {}
}

func(c *Cola[T]) Encolar(valor T) {
	c.list.InsertarUltimo(valor)
}

func(c *Cola[T]) Desencolar() (T, error) {
	if c.list.EstaVacia() {
		var zero T
		return zero, e.New("la cola está vacía")
	}
	return c.list.RemoverPrimero(), nil
}

func(c *Cola[T]) Primero() (T, error) {
	if c.list.EstaVacia() {
		var zero T
		return zero, e.New("la cola está vacía")
	}
	return c.list.verElPrimerElemento(), nil
}

func(c *Cola[T]) EstaVacia() bool {
	return c.list.EstaVacia()
}