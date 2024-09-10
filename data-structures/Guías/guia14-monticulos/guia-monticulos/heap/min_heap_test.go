package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeapCrearVacio(t *testing.T) {
	m := NewMinHeap[int]()
	assert.Equal(t, 0, m.Size())
}

func TestMinHeapRemoveMaxVacio(t *testing.T) {
	m := NewMinHeap[int]()
	_, err := m.Remove()
	assert.NotNil(t, err)
}

// Gracias a visualgo.net/en/heap
// por la ayuda para preparar este caso de prueba.
//
// Insertando los siguientes elementos en orden:
// 44, 29, 58, 2, 98, 11, 65, 3, 68, 99
//
// El arbol resultante debería ser:
//
//	[02]
//	├── [03]
//	│   ├── [29]
//	│   │   ├── [44]
//	│   │   └── [68]
//	│   └── [98]
//	│       └── [99]
//	└── [11]
//	    ├── [58]
//	    └── [65]
//
// Como arreglo:
// [02, 03, 11, 29, 98, 58, 65, 44, 68, 99].
func TestMinHeapCrearInsertarYExtraer(t *testing.T) {
	secuenciaDeInsercion := []int{44, 29, 58, 2, 98, 11, 65, 3, 68, 99}
	ordenEsperadoDespuesDeInsertar := [][]int{
		{44},
		{29, 44},
		{29, 44, 58},
		{2, 29, 58, 44},
		{2, 29, 58, 44, 98},
		{2, 29, 11, 44, 98, 58},
		{2, 29, 11, 44, 98, 58, 65},
		{2, 3, 11, 29, 98, 58, 65, 44},
		{2, 3, 11, 29, 98, 58, 65, 44, 68},
		{2, 3, 11, 29, 98, 58, 65, 44, 68, 99},
	}
	// Verificaciones iniciales
	m := NewMinHeap[int]()
	assert.Equal(t, 0, m.Size())
	// Verificaciones a medida que vamos insertando
	for i := 0; i < len(secuenciaDeInsercion); i++ {
		m.Insert(secuenciaDeInsercion[i])
		assert.Equal(t, ordenEsperadoDespuesDeInsertar[i], m.elements)
	}
	ordenEsperadoDespuesDeEliminar := [][]int{
		{3, 29, 11, 44, 98, 58, 65, 99, 68},
		{11, 29, 58, 44, 98, 68, 65, 99},
		{29, 44, 58, 99, 98, 68, 65},
		{44, 65, 58, 99, 98, 68},
		{58, 65, 68, 99, 98},
		{65, 98, 68, 99},
		{68, 98, 99},
		{98, 99},
		{99},
		{},
	}
	for i := 0; i < len(secuenciaDeInsercion); i++ {
		_, err := m.Remove()
		assert.Equal(t, ordenEsperadoDespuesDeEliminar[i], m.elements)
		assert.NoError(t, err)
	}
}

func TestNuevoMonticuloMinDesdeArregloSieteElementos(t *testing.T) {
	arreglo := []int{3, 1, 4, 1, 5, 9, 2}
	monticulo := NuevoMonticuloMinDesdeArreglo(arreglo)
	esperado := []int{1, 1, 2, 3, 5, 9, 4}
	assert.Equal(t, esperado, monticulo.elements)
}

func TestNuevoMonticuloMinDesdeArregloVacio(t *testing.T) {
	arreglo := []int{}
	monticulo := NuevoMonticuloMinDesdeArreglo(arreglo)
	assert.Equal(t, 0, monticulo.Size())
}

func TestNuevoMonticuloMinDesdeArregloQuinceElementos(t *testing.T) {
	arreglo := []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	monticulo := NuevoMonticuloMinDesdeArreglo(arreglo)
	esperado := []int{1, 6, 2, 9, 7, 5, 3, 15, 12, 13, 8, 14, 10, 11, 4}
	assert.Equal(t, esperado, monticulo.elements)
}

func TestPrimerEnesimoaMinimo(t *testing.T) {
	monticulo := NewMinHeap[int]()
	monticulo.Insert(10)
	monticulo.Insert(20)
	monticulo.Insert(30)
	monticulo.Insert(40)
	monticulo.Insert(50)
	minimo, error := monticulo.EnesimoMinimo(1)
	assert.NoError(t, error)
	assert.Equal(t, 10, minimo)
}

func TestSegundoEnesimoMinimo(t *testing.T) {
	monticulo := NewMinHeap[int]()
	monticulo.Insert(154)
	monticulo.Insert(543)
	monticulo.Insert(64)
	monticulo.Insert(7542)
	monticulo.Insert(43)
	minimo, error := monticulo.EnesimoMinimo(2)
	assert.NoError(t, error)
	assert.Equal(t, 64, minimo)
}

func TestTercerEnesimoMinimo(t *testing.T) {
	monticulo := NewMinHeap[int]()
	monticulo.Insert(53)
	monticulo.Insert(6436)
	monticulo.Insert(65)
	monticulo.Insert(85)
	monticulo.Insert(27)
	minimo, error := monticulo.EnesimoMinimo(3)
	assert.NoError(t, error)
	assert.Equal(t, 65, minimo)
}

func TestMonticulosMinimosCombinaditosVacios(t *testing.T) {
	monticulo1 := NewMinHeap[int]()
	monticulo2 := NewMinHeap[int]()
	monticulo := MonticulosCombinaditosMaximo[int](monticulo1, monticulo2)
	assert.Equal(t, 0, monticulo.Size())
}

func TestMonticulosMinimosCombinaditosUnoVacio(t *testing.T) {
	monticulo1 := NewMinHeap[int]()
	monticulo2 := NewMinHeap[int]()
	monticulo1.Insert(5)
	monticulo1.Insert(10)
	monticulo := MonticulosCombinaditosMaximo[int](monticulo1, monticulo2)
	assert.Equal(t, 2, monticulo.Size())
	elementos := monticulo.elements
	assert.Contains(t, elementos, 5)
	assert.Contains(t, elementos, 10)
}

func TestMonticulosMinimosCombinaditos(t *testing.T) {
	monticulo1 := NewMinHeap[int]()
	monticulo1.Insert(10)
	monticulo1.Insert(20)
	monticulo1.Insert(30)
	monticulo2 := NewMinHeap[int]()
	monticulo2.Insert(20)
	monticulo2.Insert(30)
	monticulo2.Insert(40)
	monticulo := MonticulosCombinaditosMaximo(monticulo1, monticulo2)
	assert.Equal(t, 6, monticulo.Size())
	elementos := monticulo.elements
	assert.Contains(t, elementos, 10)
	assert.Contains(t, elementos, 20)
	assert.Contains(t, elementos, 30)
	assert.Contains(t, elementos, 40)
}