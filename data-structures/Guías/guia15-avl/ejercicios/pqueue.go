package ejercicios_avl

import (
	"errors"
	"github.com/untref-ayp2/data-structures/types"
	"guiaAvl/avl"
)

/*
PQueue es una cola de prioridad basada en un árbol AVL.
*/
type PQueue[T types.Ordered] struct {
	tree *avl.AVLTree[T]
}

/*
NewPQueue crea una nueva PQueue vacía.
*/
func NewPQueue[T types.Ordered]() *PQueue[T] {
	return &PQueue[T]{tree: avl.NewAVLTree[T]()}
}

/*
Enqueue inserta un elemento en la cola de prioridad.
*/
func (q *PQueue[T]) Enqueue(v T) {
	q.tree.Insert(v)
}

/*
Dequeue elimina y devuelve el elemento de mayor prioridad (el más grande) de la cola.
*/
func (q *PQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var emptyValue T
		return emptyValue, errors.New("cola vacía")
	}
	max, _ := q.tree.FindMax()
	q.tree.Remove(max)
	return max, nil
}

/*
Front devuelve el elemento de mayor prioridad (el más grande) sin eliminarlo de la cola.
*/
func (q *PQueue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var emptyValue T
		return emptyValue, errors.New("cola vacía")
	}
	max, _ := q.tree.FindMax()
	return max, nil
}

/*
IsEmpty verifica si la cola de prioridad está vacía.
*/
func (q *PQueue[T]) IsEmpty() bool {
	return q.tree.IsEmpty()
}