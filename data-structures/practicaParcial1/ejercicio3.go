package practicaParcial1

//Dada una pila y una cola de N números enteros, generar una lista simplemente enlazada, insertando los elementos en forma intercalada, comenzando por la pila y luego por la cola.

import(
	s "github.com/trigologiaa/data-structures/stack"
	q "github.com/trigologiaa/data-structures/queue"
	l "github.com/trigologiaa/data-structures/list"
)

func IntercalarPilaYCola(pila *s.Stack[int], cola *q.Queue[int]) *l.LinkedList[int] {
	lista := l.NewLinkedList[int]()
	for !pila.IsEmpty() || !cola.IsEmpty() {
		if !pila.IsEmpty() {
			valor, err := pila.Pop()
			if err == nil {
				lista.Append(valor)
			}
		}
		if !cola.IsEmpty() {
			valor, err := cola.Dequeue()
			if err == nil {
				lista.Append(valor)
			}
		}
	}
	return lista
}