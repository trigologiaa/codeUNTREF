package practicaParcial2

//Escribir un método recursivo que reciba un árbol binario de búsqueda cuyas claves son enteros y devuelva la suma de sus hojas cuyos valores son < 10. Si el árbol está vacío debe devolver cero.

import(
	bst "github.com/trigologiaa/data-structures/tree/binarytree"
)

func SumaDeHojasMenoresQue10(bst *bst.BinarySearchTree[int]) int {
	if bst.Size() == 0 {
		return 0
	}
	return sumaDeHojasMenoresQue10(bst.GetRoot())
}

func sumaDeHojasMenoresQue10(n *bst.BinaryNode[int]) int {
	if n == nil {
		return 0
	}
	if n.GetLeft() == nil && n.GetRight() == nil {
		if n.GetData() < 10 {
			return n.GetData()
		} else {
			return 0
		}
	}
	return sumaDeHojasMenoresQue10(n.GetLeft()) + sumaDeHojasMenoresQue10(n.GetRight())
}