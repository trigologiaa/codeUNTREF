package practicaParcial1

//Implementar un conjunto sobre una lista enlazada simple

import (
	f "fmt"
	l "github.com/trigologiaa/data-structures/list"
)

type ListSet[T comparable] struct {
	elementos *l.LinkedList[T]
}

func NewListSet[T comparable](elementos ...T) *ListSet[T] {
	set := &ListSet[T] {
		l.NewLinkedList[T](),
	}
	set.Add(elementos...)
	return set
}

func (s *ListSet[T]) Contains(elemento T) bool {
	return s.elementos.Find(elemento) != nil
}

func (s *ListSet[T]) Add(elementos ...T) {
	for _, elemento := range elementos {
		if !s.Contains(elemento) {
			s.elementos.Append(elemento)
		}
	}
}

func (s *ListSet[T]) Remove(elemento T) {
	s.elementos.Remove(elemento)
}

func (s *ListSet[T]) Size() int {
	return s.elementos.Size()
}

func (s *ListSet[T]) Values() []T {
	var valores []T
	for nodo := s.elementos.Head(); nodo != nil; nodo = nodo.Next() {
		valores = append(valores, nodo.Data())
	}
	return valores
}

func (s *ListSet[T]) String() string {
	str := "Set: {"
	for indice, valor := range s.Values() {
		if indice > 0 {
			str += ", "
		}
		str += f.Sprintf("%v", valor)
	}
	str += "}"
	return str
}