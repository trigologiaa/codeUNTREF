package avl

import (
	"fmt"
	"github.com/untref-ayp2/data-structures/types"
	"github.com/untref-ayp2/data-structures/utils"
)

// AVLNode representa un nodo en un árbol AVL.
//
//
//
//
//
//
//
//
type AVLNode[T types.Ordered] struct {
	data   T           // Dato
	height int         // Altura
	left   *AVLNode[T] // Hijo izquierdo del nodo
	right  *AVLNode[T] // Hijo derecho del nodo
}

// newAVLNode crea un nuevo nodo AVL con el valor dado y los hijos especificados.
//
//
//
//
//
//
//
//
func newAVLNode[T types.Ordered](data T, left *AVLNode[T], right *AVLNode[T]) *AVLNode[T] {
	return &AVLNode[T]{left: left, right: right, data: data, height: 0}
}

// GetData devuelve el dato almacenado en el nodo.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) GetData() T {
	return n.data
}

// String devuelve una representación en string del dato almacenado en el nodo.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) string() string {
	return fmt.Sprintf("%v", n.data)
}

// GetLeft devuelve el hijo izquierdo del nodo.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) GetLeft() *AVLNode[T] {
	return n.left
}

// GetRight devuelve el hijo derecho del nodo.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) GetRight() *AVLNode[T] {
	return n.right
}

// GetHeight devuelve la altura del nodo en el árbol.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) GetHeight() int {
	if n == nil {
		return -1
	}
	return n.height
}

// GetBalance calcula el factor de balance del nodo.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) GetBalance() int {
	if n == nil {
		return 0
	}
	return n.left.GetHeight() - n.right.GetHeight()
}

// updateHeight actualiza la altura del nodo basada en las alturas de sus hijos.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) updateHeight() {
	n.height = 1 + utils.Max(n.left.GetHeight(), n.right.GetHeight())
}

// insert inserta un nuevo valor en el árbol AVL.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) insert(value T) *AVLNode[T] {
	//Si el árbol está vacío
	if n == nil {
		//Se crea un nuevo árbol con 'value' como root y los hijos izquierdo y derecho con valor nulo
		return newAVLNode[T](value, nil, nil)
	}
	//Casos donde el árbol no está vacío
	switch {
	//Cuando el valor insertado sea menor que el nodo root
	case value < n.data:
		//Se actualiza el valor del hijo izquierdo al mismo del 'value' insertado
		n.left = n.left.insert(value)
	//Cuando el valor insertado sea mayor que el nodo root
	case value > n.data:
		//Se actualiza el valor del hijo derecho al mismo del 'value' insertado
		n.right = n.right.insert(value)
	default:
		//Se retorna el valor del árbol
		return n
	}
	//Se actualiza la altura del árbol
	n.updateHeight()
	//Se retorna el árbol con la rotación aplicada
	return n.applyRotation()
}

// rotateRight realiza una rotación a la derecha en el nodo.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) rotateRight() *AVLNode[T] {
	y := n.left
	t2 := y.right
	y.right = n
	n.left = t2
	n.updateHeight()
	y.updateHeight()
	return y
}

// rotateLeft realiza una rotación a la izquierda en el nodo.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) rotateLeft() *AVLNode[T] {
	//Se crea la variable 'x' de tipo *AVLNode[T] y se le asigna el valor del hijo derecho del nodo 'n'. 'x' es un nodo.
	x := n.right
	//Se crea la variable 't2' de tipo *AVLNode[T] y se le asigna el valor del hijo izquierdo del nodo 'x'. 't2' es un nodo.
	t2 := x.left
	//Al hijo izquierdo del nodo 'x' se le asigna el valor del nodo 'n'.
	x.left = n
	//Al hijo derecho del nodo 'n' se le asigna el valor del nodo 't2'.
	n.right = t2
	//Se actualiza la altura del nodo 'n'.
	n.updateHeight()
	//Se actualiza la altura del nodo 'x'.
	x.updateHeight()
	//Se retorna el valor final del nodo 'x'.
	return x
}

// remove elimina un valor del árbol AVL.
//
//
//
//
//
//
//
//
func (n *AVLNode[T]) remove(value T) *AVLNode[T] {
	if n == nil {
		return n
	}
	switch {
	case value < n.data:
		n.left = n.left.remove(value)
	case value > n.data:
		n.right = n.right.remove(value)
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		temp := n.right.findMin()
		n.data = temp.data
		n.right = n.right.remove(temp.data)
	}
	n.updateHeight()
	return n.applyRotation()
}

// applyRotation verifica el balance del nodo actual y aplica las rotaciones necesarias para mantener el equilibrio del árbol AVL.
//
// Uso:
//		tree := NewAVLTree[int]()
//		rotatedNode := tree.root.applyRotation
//
// Retorna:
//		- 
//
func (n *AVLNode[T]) applyRotation() *AVLNode[T] {
	//Se crea la variable 'balance' que calcula el factor de balance del nodo.
	balance := n.GetBalance()
	//Si el factor del balance es mayor a 1:
	if balance > 1 {
		//Si el factor del balance del hijo izquierdo del nodo es menor a 0:
		if n.left.GetBalance() < 0 {
			//Se retorna una rotación izquierda en el hijo izquierdo.
			n.left = n.left.rotateLeft()
		}
		//Se retorna una rotación derecha en el nodo actual.
		return n.rotateRight()
	}
	//Si el factor de balance es menor a -1:
	if balance < -1 {
		//Si el factor
		if n.right.GetBalance() > 0 {
			n.right = n.right.rotateRight()
		}
		//Se retorna una rotación izquierda en el nodo actual.
		return n.rotateLeft()
	}
	//Se retorna el nodo actual.
	return n
}

// findMin busca el valor mínimo en un subárbol con raíz en el nodo actual.
//
// Uso:
//		tree := NewAVLTree[int]()
//		minNode := tree.root.findMin()
//
// Retorna:
//		- El valor mínimo en el subárbol como un puntero al nodo 'n'.
func (n *AVLNode[T]) findMin() *AVLNode[T] {
	//Si el hijo izquierdo es nulo:
	if n.left == nil {
		//Se retorna el nodo actual.
		return n
	}
	//Se retorna un llamado recursivo a findMin() en el hijo izquierdo del nodo.
	return n.left.findMin()
}

// findMax busca el valor máximo en un subárbol con raíz en el nodo actual.
//
// Uso:
//		tree := NewALVTree[int]()
//		maxNode := tree.root.findMax()
//
// Retorna:
//		- El valor máximo en el subárbol como un puntero al nodo 'n'.
func (n *AVLNode[T]) findMax() *AVLNode[T] {
	//Si el hijo derecho es nulo:
	if n.right == nil {
		//Se retorna el nodo actual.
		return n
	}
	//Se retorna un llamado recursivo a findMax() en el hijo derecho del nodo.
	return n.right.findMax()
}

// search busca un valor en el árbol AVL.
//
// Parámetros:
//		- k: valor T genérico que se desea buscar en el subárbol.
//
// Uso:
// 		tree := NewAVLTree[int]()
// 		found := tree.root.search(5)
//
// Retorna:
//		- 'true' si el valor 'k' se encuentra en el subárbol como un bool.
//		- 'false' si el valor 'k' no se encuentra en el subárbol como un bool.
func (n *AVLNode[T]) search(k T) bool {
	//Si el nodo es nulo:
	if n == nil {
		//Se retorna false.
		return false
	}
	//Si el nodo es mayor al valor buscado:
	if n.data > k {
		//Se retorna un llamado recursivo a search(k) en el hijo izquierdo del nodo.
		return n.left.search(k)
	}
	//Si el nodo es menor al valor buscado:
	if n.data < k {
		//Se retorna un llamado recursivo a search(k) en el hijo derecho del nodo.
		return n.right.search(k)
	}
	//Se retorna true ya que el dato del nodo es igual al valor buscado.
	return true
}

// inOrder devuelve una representación en orden del árbol como una cadena.
//		/Si el nodo es nulo:
//			/Se retorna un string vacío.
//		/Si no, se retorna una representación en orden del subárbol.
//
//
//
//
//
func (n *AVLNode[T]) inOrder() string {
	if n == nil {
		return ""
	}
	return n.left.inOrder() + " " + n.string() + " " + n.right.inOrder()
}

/*
Size devuelve el número total de nodos en el subárbol.
		/Si el nodo es nulo:
			/Se retorna 0.
		/Si no, se retorna el tamaño del subárbol sumando 1 más el tamaño de los subárboles izquierdo y derecho.
Uso:
 	tree := NewTreeSet[int]()
 	size := tree.root.Size()
Retorna:
	- El número total de nodos en el subárbol como un int.
*/
func (n *AVLNode[T]) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}