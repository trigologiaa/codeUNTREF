// Package heap provee una implementación de un heap binario.
package heap

import (
	"errors"
	"github.com/untref-ayp2/data-structures/types"
	"github.com/untref-ayp2/data-structures/utils"
)

type Heap[T any] struct {
	// contenedor de datos
	elements []T
	// Función de comparación. Para un heap de mínimo,
	// devuelve -1 si a < b, 0 si a == b, 1 si a > b
	// Para un heap de máximo, devuelve 1 si a < b, 0 si a == b, -1 si a > b
	compare func(a T, b T) int
}

// NewMinHeap crea un nuevo heap binario de mínimos.
//
// Uso:
//
//	heap := heap.NewMinHeap[int]()
//
// Retorna:
//   - un puntero a un heap binario de mínimos.
func NewMinHeap[T types.Ordered]() *Heap[T] {
	return &Heap[T]{compare: utils.Compare[T], elements: make([]T, 0)}
}

// NewMaxHeap crea un nuevo heap binario de máximos.
//
// Uso:
//
//	heap := heap.NewMaxHeap[int]()
//
// Retorna:
//   - un puntero a un heap binario de máximos.
func NewMaxHeap[T types.Ordered]() *Heap[T] {
	comp := func(a T, b T) int {
		return utils.Compare[T](b, a)
	}
	return &Heap[T]{compare: comp, elements: make([]T, 0)}
}

// NewGenericHeap crea un nuevo heap binario con una función de comparación personalizada.
//
// Uso:
//
//	heap := heap.NewGenericHeap[int](func(a int, b int) int {
//		if a < b {
//			return -1
//		}
//		if a > b {
//			return 1
//		}
//		return 0
//	})
//
// Parámetros:
//   - `comp` función de comparación personalizada.
//
// Retorna:
//   - un puntero a un heap binario con una función de comparación personalizada.
func NewGenericHeap[T any](comp func(a T, b T) int) *Heap[T] {
	return &Heap[T]{compare: comp, elements: make([]T, 0)}
}

// Size retorna la cantidad de elementos en el heap.
//
// Uso:
//
//	size := heap.Size()
//
// Retorna:
//   - la cantidad de elementos en el heap.
func (m *Heap[T]) Size() int {
	return len(m.elements)
}

// Insert agrega un elemento al heap.
//
// Uso:
//
//	heap := heap.NewMinHeap[int]()
//	heap.Insert(5)
//
// Parámetros:
//
//	element: elemento a agregar al heap.
func (m *Heap[T]) Insert(element T) {
	m.elements = append(m.elements, element)
	m.upHeap(len(m.elements) - 1)
}

// upHeap reordena el heap hacia arriba.
//
// Parámetros:
//   - `i` índice del elemento a reordenar.
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

// Remove elimina y retorna el elemento en la cima del heap.
//
// Uso:
//
//	heap := heap.NewMinHeap[int]()
//	heap.Insert(5)
//	element, _ := heap.Remove()
//
// Retorna:
//   - el elemento en la cima del heap.
func (m *Heap[T]) Remove() (T, error) {
	var element T
	if m.Size() == 0 {
		return element, errors.New("heap vacío")
	}
	element = m.elements[0]
	m.elements[0] = m.elements[m.Size()-1]
	m.elements = m.elements[:m.Size()-1]
	m.downHeap(0)
	return element, nil
}

// downHeap reordena el heap hacia abajo.
//
// Parámetros:
//   - `i` índice del elemento a reordenar.
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

// NuevoMonticuloMaxDesdeArreglo crea un nuevo montículo de máximo a partir de un arreglo desordenado.
//
// Uso:
//
//	arreglo := []int{3, 1, 4, 1, 5, 9, 2}
//	heap := NuevoMonticuloMaxDesdeArreglo(arreglo)
//
// Parámetros:
//   - `arreglo`: arreglo desordenado de elementos.
//
// Retorna:
//   - un puntero a un montículo de máximo que contiene los elementos del arreglo.
func NuevoMonticuloMaxDesdeArreglo[T types.Ordered](arreglo []T) *Heap[T] {
	monticulo := &Heap[T]{compare: func(a, b T) int {return utils.Compare[T](b, a)}, elements: arreglo}
	for i := len(arreglo) / 2 - 1; i >= 0; i-- {
		monticulo.downHeap(i)
	}
	return monticulo
}

// NuevoMonticuloMinDesdeArreglo crea un nuevo montículo de mínimo a partir de un arreglo desordenado de elementos.
//
// Uso:
//
//	arreglo := []int{5, 2, 8, 1, 9}
//	heap := NuevoMonticuloMinDesdeArreglo(arreglo)
//
// Parámetros:
//   - arreglo: el arreglo desordenado de elementos.
//
// Retorna:
//   - un nuevo montículo de mínimo que contiene los elementos del arreglo.
func NuevoMonticuloMinDesdeArreglo[T types.Ordered](arreglo []T) *Heap[T] {
	monticulo := NewMinHeap[T]()
	for _, elemento := range arreglo {
		monticulo.Insert(elemento)
	}
	return monticulo
}

// EnesimoMaximo devuelve el enésimo máximo del montículo.
//
// Parámetros:
//   - n: índice del máximo deseado (1 para el máximo, 2 para el segundo máximo, etc.).
//
// Retorna:
//   - el enésimo máximo del montículo.
//   - un error si el índice n está fuera de rango.
func (m *Heap[T]) EnesimoMaximo(n int) (T, error) {
	if n < 1 || n > m.Size() {
		var cero T
		return cero, errors.New("índice fuera de rango")
	}
	copia := make([]T, len(m.elements))
	copy(copia, m.elements)
	for i := 0; i < n-1; i++ {
		_, error := m.Remove()
		if error != nil {
			var cero T
			return cero, error
		}
	}
	maximo, _ := m.Remove()
	m.elements = copia
	return maximo, nil
}

func (m *Heap[T]) EnesimoMinimo(n int) (T, error) {
	if n < 1 || n > m.Size() {
		var cero T
		return cero, errors.New("índice fuera de rango")
	}
	copia := make([]T, len(m.elements))
	copy(copia, m.elements)
	for i := 0; i < n-1; i++ {
		_, error := m.Remove()
		if error != nil {
			var cero T
			return cero, error
		}
	}
	minimo, _ := m.Remove()
	m.elements = copia
	return minimo, nil
}

// MonticulosCombinaditos combina dos montículos y devuelve un tercero que es la combinación de ambos.
//
// Parámetros:
//   - heap1: el primer montículo.
//   - heap2: el segundo montículo.
//
// Retorna:
//   - un nuevo montículo que contiene todos los elementos de heap1 y heap2.
func MonticulosCombinaditosMaximo[T any](heap1, heap2 *Heap[T]) *Heap[T] {
	monticuloCombinado := NewGenericHeap[T](heap1.compare)
	for _, elemento := range heap1.elements {
		monticuloCombinado.Insert(elemento)
	}
	for _, elemento := range heap2.elements {
		monticuloCombinado.Insert(elemento)
	}
	return monticuloCombinado
}

func MonticulosCombinaditosMinimo[T any](heap1, heap2 *Heap[T]) *Heap[T] {
	monticuloCombinado := NewGenericHeap[T](heap1.compare)
	for _, elemento := range heap1.elements {
		monticuloCombinado.Insert(elemento)
	}
	for _, elemento := range heap2.elements {
		monticuloCombinado.Insert(elemento)
	}
	return monticuloCombinado
}