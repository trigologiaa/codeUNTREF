package practicaParcial2

//Describir para que se usan los métodos UpHeap y DownHeap en un montículo, escribir dichos métodos en GO y mostrar paso a paso el resultado de realizar las siguientes operaciones en un heap de mínimos (de enteros) recién creado: insertar 10, insertar 4, insertar 47, insertar 1, insertar 27, insertar -5, insertar 2, borrar mínimo, borrar mínimo, insertar 7, insertar -1.

/*
UpHeap se utiliza para reordenar el heap hacia arriba, esto después de confirmar que el último insertado es menor o mayor que su padre (dependiendo de si es heap de mínimo o heap de máximo respectivamente).
DownHeap se utiliza cuando se elimina el nodo raíz, entonces el último elemento se coloca sobre la raíz del heap y se compara con sus hijos, cambiando de lugar con el hijo que tenga el valor más bajo o alto (dependiendo de si es un heap de mínimo o un heap de máximo respectivamente), y así repitiendo el proceso hasta que el nodo que se está comparando deje de cumplir esa condición.
*/

/*
func (m *Heap[T]) upHeap(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if m.compare(m.elements[i], m.elements[parent]) > 0 {
			break
		}
		m.elements[i], m.elements[parent] = m.elements[parent], m.elements[i]
		i = parent
	}
}

func (m *Heap[T]) downHeap(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i
		if left < m.Size() && m.compare(m.elements[left], m.elements[smallest]) < 0 {
			smallest = left
		}
		if right < m.Size() && m.compare(m.elements[right], m.elements[smallest]) < 0 {
			smallest = right
		}
		if smallest == i {
			break
		}
		m.elements[i], m.elements[smallest] = m.elements[smallest], m.elements[i]
		i = smallest
	}
}
	*/