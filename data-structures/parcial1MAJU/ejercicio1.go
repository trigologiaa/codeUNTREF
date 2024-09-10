package parcial1MAJU

/*
Dada la implementación de Set vista en clase (implementado sobre una lista enlazada simple) que tiene los métodos: Contains, Add, Remove, Size, String y Values y soporta las siguientes operaciones que en todos los casos devuelven un conjunto nuevo: NewSet, Union, Intersection, Difference, Subset y Equal. Se pide reescribir la operación Union para que reciba como parámetro un número indefinido de conjuntos y devuelva la Union de todos ellos.
func Union[T comparable](conjuntos ...*Set[T]) *Set[T]
*/

import(
	ls "github.com/trigologiaa/data-structures/set"
)

func Union[T comparable](conjuntos ...*ls.ListSet[T]) *ls.ListSet[T] {
	resultado := ls.NewListSet[T]()
	for _, conjunto := range conjuntos {
		for _, elemento := range conjunto.Values() {
			resultado.Add(elemento)
		}
	}
	return resultado
}