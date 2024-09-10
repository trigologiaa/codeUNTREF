package heap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMaxHeapCrearVacio(t *testing.T) {
	m := NewMaxHeap[int]()
	assert.Equal(t, 0, m.Size())
}

func TestMaxHeapRemoveMaxVacio(t *testing.T) {
	m := NewMaxHeap[int]()
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
//	[99]
//	├── [98]
//	│   ├── [58]
//	│   │   ├── [2]
//	│   │   └── [3]
//	│   └── [68]
//	│       └── [29]
//	└── [65]
//	    ├── [11]
//	    └── [44]
//
// Como arreglo:
// [99, 98, 65, 58, 68, 11, 44, 2, 3, 29].
func TestMaxHeapCrearInsertarYExtraer(t *testing.T) {
	secuenciaDeInsercion := []int{44, 29, 58, 2, 98, 11, 65, 3, 68, 99}
	ordenEsperadoDespuesDeInsertar := [][]int{
		{44},
		{44, 29},
		{58, 29, 44},
		{58, 29, 44, 2},
		{98, 58, 44, 2, 29},
		{98, 58, 44, 2, 29, 11},
		{98, 58, 65, 2, 29, 11, 44},
		{98, 58, 65, 3, 29, 11, 44, 2},
		{98, 68, 65, 58, 29, 11, 44, 2, 3},
		{99, 98, 65, 58, 68, 11, 44, 2, 3, 29},
	}
	// Verificaciones iniciales
	m := NewMaxHeap[int]()
	assert.Equal(t, 0, m.Size())
	// Verificaciones a medida que vamos insertando
	for i := 0; i < len(secuenciaDeInsercion); i++ {
		m.Insert(secuenciaDeInsercion[i])
		assert.Equal(t, ordenEsperadoDespuesDeInsertar[i], m.elements)
	}
	ordenEsperadoDespuesDeEliminar := [][]int{
		{98, 68, 65, 58, 29, 11, 44, 2, 3},
		{68, 58, 65, 3, 29, 11, 44, 2},
		{65, 58, 44, 3, 29, 11, 2},
		{58, 29, 44, 3, 2, 11},
		{44, 29, 11, 3, 2},
		{29, 3, 11, 2},
		{11, 3, 2},
		{3, 2},
		{2},
		{},
	}
	for i := 0; i < len(secuenciaDeInsercion); i++ {
		_, err := m.Remove()
		assert.Equal(t, ordenEsperadoDespuesDeEliminar[i], m.elements)
		assert.NoError(t, err)
	}
}

func TestNuevoMonticuloMaxDesdeArregloSieteElementos(t *testing.T) {
	arreglo := []int{3, 1, 4, 1, 5, 9, 2}
	monticulo := NuevoMonticuloMaxDesdeArreglo(arreglo)
	esperado := []int{9, 5, 4, 1, 1, 3, 2}
	assert.Equal(t, esperado, monticulo.elements)
}

func TestNuevoMonticuloMaxDesdeArregloVacio(t *testing.T) {
	arreglo := []int{}
	monticulo := NuevoMonticuloMaxDesdeArreglo(arreglo)
	assert.Equal(t, 0, monticulo.Size())
}

func TestNuevoMonticuloMaxDesdeArregloQuinceElementos(t *testing.T) {
	arreglo := []int{2, 7, 26, 25, 19, 17, 1, 90, 3, 36, 12, 4, 6, 14, 87}
	monticulo := NuevoMonticuloMaxDesdeArreglo(arreglo)
	esperado := []int{90, 36, 87, 25, 19, 17, 26, 7, 3, 2, 12, 4, 6, 14, 1}
	assert.Equal(t, esperado, monticulo.elements)
}

func TestPrimerEnesimoMaximo(t *testing.T) {
	monticulo := NewMaxHeap[int]()
	monticulo.Insert(10)
	monticulo.Insert(20)
	monticulo.Insert(30)
	monticulo.Insert(40)
	monticulo.Insert(50)
	maximo, error := monticulo.EnesimoMaximo(1)
	assert.NoError(t, error)
	assert.Equal(t, 50, maximo)
}

func TestSegundoEnesimoMaximo(t *testing.T) {
	monticulo := NewMaxHeap[int]()
	monticulo.Insert(154)
	monticulo.Insert(543)
	monticulo.Insert(64)
	monticulo.Insert(7542)
	monticulo.Insert(43)
	maximo, error := monticulo.EnesimoMaximo(2)
	assert.NoError(t, error)
	assert.Equal(t, 543, maximo)
}

func TestTercerEnesimoMaximo(t *testing.T) {
	monticulo := NewMaxHeap[int]()
	monticulo.Insert(53)
	monticulo.Insert(6436)
	monticulo.Insert(65)
	monticulo.Insert(85)
	monticulo.Insert(27)
	maximo, error := monticulo.EnesimoMaximo(3)
	assert.NoError(t, error)
	assert.Equal(t, 65, maximo)
}

func TestMonticulosMaximosCombinaditosVacios(t *testing.T) {
	monticulo1 := NewMaxHeap[int]()
	monticulo2 := NewMaxHeap[int]()
	monticulo := MonticulosCombinaditosMaximo[int](monticulo1, monticulo2)
	assert.Equal(t, 0, monticulo.Size())
}

func TestMonticulosMaximosCombinaditosUnoVacio(t *testing.T) {
	monticulo1 := NewMaxHeap[int]()
	monticulo2 := NewMaxHeap[int]()
	monticulo1.Insert(5)
	monticulo1.Insert(10)
	monticulo := MonticulosCombinaditosMaximo[int](monticulo1, monticulo2)
	assert.Equal(t, 2, monticulo.Size())
	elementos := monticulo.elements
	assert.Contains(t, elementos, 5)
	assert.Contains(t, elementos, 10)
}

func TestMonticulosMaximosCombinaditos(t *testing.T) {
	monticulo1 := NewMaxHeap[int]()
	monticulo1.Insert(10)
	monticulo1.Insert(20)
	monticulo1.Insert(30)
	monticulo2 := NewMaxHeap[int]()
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