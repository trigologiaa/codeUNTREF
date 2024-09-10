package ejercicios_avl

import (
	"fmt"
	"strings"
	"github.com/untref-ayp2/data-structures/types"
	"guiaAvl/avl"
)

/*
TreeSet es un conjunto basado en un árbol AVL.
*/
type TreeSet[T types.Ordered] struct {
	tree *avl.AVLTree[T]
}

/*
NewTreeSet crea un nuevo TreeSet e inserta los elementos opcionales dados.
*/
func NewTreeSet[T types.Ordered](elements ...T) *TreeSet[T] {
	set := &TreeSet[T]{tree: avl.NewAVLTree[T]()}
	for _, elem := range elements {
		set.Add(elem)
	}
	return set
}

/*
Contains verifica si un elemento está presente en el conjunto.
*/
func (s *TreeSet[T]) Contains(element T) bool {
	return s.tree.Search(element)
}

/*
Add inserta uno o más elementos en el conjunto.
*/
func (s *TreeSet[T]) Add(elements ...T) {
	for _, elem := range elements {
		s.tree.Insert(elem)
	}
}


/*
Remove elimina un elemento del conjunto.
*/
func (s *TreeSet[T]) Remove(element T) {
	s.tree.Remove(element)
}


/*
Size devuelve el número de elementos en el conjunto.
*/
func (s *TreeSet[T]) Size() int {
	return s.tree.Size()
}


/*
Values devuelve una lista de los elementos en el conjunto en orden ascendente.
*/
func (s *TreeSet[T]) Values() []T {
	var values []T
	it := avl.NewAVLInOrderIterator(s.tree.GetRoot())
	for it.HasNext() {
		value, _ := it.Next()
		values = append(values, value)
	}
	return values
}


/*
String devuelve una representación en cadena del conjunto.
*/
func (s *TreeSet[T]) String() string {
	values := s.Values()
	valueStrings := make([]string, len(values))
	for i, v := range values {
		valueStrings[i] = fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("Set: {%s}", strings.Join(valueStrings, ", "))
}