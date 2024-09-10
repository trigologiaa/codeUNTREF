package ejercicios_avl

import (
	"github.com/untref-ayp2/data-structures/types"
	"guiaAvl/avl"
)

func AVLTreeSort[T types.Ordered](elements []T) []T {
	tree := avl.NewAVLTree[T]()
	for _, elem := range elements {
		tree.Insert(elem)
	}
	sorted := []T{}
	iterator := avl.NewAVLInOrderIterator(tree.GetRoot())
	for iterator.HasNext() {
		elem, _ := iterator.Next()
		sorted = append(sorted, elem)
	}
	return sorted
}