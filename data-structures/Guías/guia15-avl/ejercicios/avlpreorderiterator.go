package ejercicios_avl

import (
	"guiaAvl/avl"
	"github.com/untref-ayp2/data-structures/stack"
	"github.com/untref-ayp2/data-structures/types"
	"errors"
)

/*
AVLPreOrderIterator Define el iterador preorden para el árbol AVL.
*/
type AVLPreOrderIterator[T types.Ordered] struct {
	stack *stack.Stack[*avl.AVLNode[T]]
}

/*
NewAVLPreOrderIterator crea un nuevo iterador preorden para un árbol AVL a partir de la raíz.
*/
func NewAVLPreOrderIterator[T types.Ordered](root *avl.AVLNode[T]) *AVLPreOrderIterator[T] {
	iterator := &AVLPreOrderIterator[T]{
		stack: stack.NewStack[*avl.AVLNode[T]](),
	}
	if root != nil {
		iterator.stack.Push(root)
	}
	return iterator
}
/*
HasNext verifica si hay más nodos para visitar en el recorrido preorden.
*/
func (it *AVLPreOrderIterator[T]) HasNext() bool {
	return !it.stack.IsEmpty()
}

/*
Next devuelve el siguiente valor en el recorrido preorden.
*/
func (it *AVLPreOrderIterator[T]) Next() (T, error) {
	var data T
	if it.stack.IsEmpty() {
		return data, errors.New("no hay más elementos")
	}
	currentNode, _ := it.stack.Pop()
	data = currentNode.GetData()
	if currentNode.GetRight() != nil {
		it.stack.Push(currentNode.GetRight())
	}
	if currentNode.GetLeft() != nil {
		it.stack.Push(currentNode.GetLeft())
	}
	return data, nil
}