package parcial1MAJU2TA

//Escribir una función recursiva que reciba como parámetro una cola y devuelva la cola invertida, es decir, el último elemento de la cola pasa a ser el primero y viceversa. No utilizar ninguna estructura de datos auxiliar.

import(
	c "github.com/trigologiaa/data-structures/queue"
)

func Invertir[T any](cola *c.Queue[T]) *c.Queue[T] {
	if cola.IsEmpty() {
		return cola
	}
	elemento, _ := cola.Dequeue()
	colaInvertida := Invertir[T](cola)
	colaInvertida.Enqueue(elemento)
	return colaInvertida
}