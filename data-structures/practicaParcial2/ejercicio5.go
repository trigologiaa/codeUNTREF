package practicaParcial2

//Escribir un método recursivo que reciba un árbol binario de enteros (no de búsqueda) y reemplace el valor de cada nodo con la suma de sus hijos.
//Imagen ilustrativa.

//Se crea el método SetData(data T) para BinaryNode.

import(
	bt "github.com/trigologiaa/data-structures/tree/binarytree"
	t "github.com/trigologiaa/data-structures/types"
)

func SumaDeHijos[T t.Ordered](n *bt.BinaryNode[T]) T {
	if n == nil {
		var zero T
		return zero
	}
	sumaIzquierdo := SumaDeHijos[T](n.GetLeft())
	sumaDerecho := SumaDeHijos[T](n.GetRight())
	valorActual := n.GetData()
	n.SetData(sumaIzquierdo + sumaDerecho)
	return valorActual + sumaIzquierdo + sumaDerecho
}

func SumaDeHijosWrapper[T t.Ordered](t *bt.BinaryTree[T]) {
	if t == nil || t.GetRoot() == nil {
		return
	}
	SumaDeHijos[T](t.GetRoot())
}