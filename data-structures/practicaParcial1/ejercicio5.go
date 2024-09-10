package practicaParcial1

//Escribir una función que reciba como parámetro dos listas enlazadas simples del mismo tipo y devuelva la lista que resulta de intercalar uno a uno cada uno de los elementos de las listas que recibió como parámetro.

import(
	l "github.com/trigologiaa/data-structures/list"
)

func Intercalar[T comparable](lista1, lista2 *l.LinkedList[T]) *l.LinkedList[T] {
	resultado := l.NewLinkedList[T]()
	actual1 := lista1.Head()
	actual2 := lista2.Head()
	for actual1 != nil && actual2 != nil {
		resultado.Append(actual1.Data())
		resultado.Append(actual2.Data())
		actual1 = actual1.Next()
		actual2 = actual2.Next()
	}
	for actual1 != nil {
		resultado.Append(actual1.Data())
		actual1 = actual1.Next()
	}
	for actual2 != nil {
		resultado.Append(actual2.Data())
		actual2 = actual2.Next()
	}
	return resultado
}