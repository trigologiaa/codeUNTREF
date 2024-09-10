package parcial1MAMI

//Escribir una función recursiva que reciba como parámetro una Cola y devuelva una Cola invertida. Por ejemplo, si la Cola que recibe es [1, 2, 3, 4, 5] donde 1 es el primer elemento que entró en la cola, debe devolver [5, 4, 3, 2, 1]

import(
	q "github.com/trigologiaa/data-structures/queue"
)

func InvertirCola[T any](q *q.Queue[T]) *q.Queue[T] {
	if q.IsEmpty() {
		return q
	}
	elementoDelFrente, _ := q.Dequeue()
	invertida := InvertirCola(q)
	invertida.Enqueue(elementoDelFrente)
	return invertida
}