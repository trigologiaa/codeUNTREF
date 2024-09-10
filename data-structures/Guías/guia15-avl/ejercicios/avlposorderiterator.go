package ejercicios_avl

import (
	"errors"
	"guiaAvl/avl"
	"github.com/untref-ayp2/data-structures/stack"
	"github.com/untref-ayp2/data-structures/types"
)

/*
AVLPosOrderIterator define el iterador postorden para el árbol AVL.
*/
type AVLPosOrderIterator[T types.Ordered] struct {
	stack1 *stack.Stack[*avl.AVLNode[T]]
	stack2 *stack.Stack[*avl.AVLNode[T]]
}

/*
NewAVLPosOrderIterator crea un nuevo iterador postorden para un árbol AVL a partir de la raíz.
*/
func NewAVLPosOrderIterator[T types.Ordered](root *avl.AVLNode[T]) *AVLPosOrderIterator[T] {
	iterator := &AVLPosOrderIterator[T]{
		stack1: stack.NewStack[*avl.AVLNode[T]](),
		stack2: stack.NewStack[*avl.AVLNode[T]](),
	}
	if root != nil {
		iterator.stack1.Push(root)
		for !iterator.stack1.IsEmpty() {
			node, _ := iterator.stack1.Pop()
			iterator.stack2.Push(node)
			if node.GetLeft() != nil {
				iterator.stack1.Push(node.GetLeft())
			}
			if node.GetRight() != nil {
				iterator.stack1.Push(node.GetRight())
			}
		}
	}
	return iterator
}

/*
HasNext verifica si hay más nodos para visitar en el recorrido postorden.
*/
func (it *AVLPosOrderIterator[T]) HasNext() bool {
	return !it.stack2.IsEmpty()
}

/*
Next devuelve el siguiente valor en el recorrido postorden.
*/
func (it *AVLPosOrderIterator[T]) Next() (T, error) {
	var data T
	if it.stack2.IsEmpty() {
		return data, errors.New("no hay más elementos")
	}
	node, _ := it.stack2.Pop()
	return node.GetData(), nil
}