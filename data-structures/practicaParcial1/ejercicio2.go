package practicaParcial1

//Dada una lista de números, escriba una función para encontrar el número máximo utilizando una pila.

import(
	e "errors"
	s "github.com/trigologiaa/data-structures/stack"
)

func EncontrarMaximo(lista []int) (int, error) {
	if len(lista) == 0 {
		return 0, e.New("la lista está vacía")
	}
	pila := s.NewStack[int]()
	for _, numero := range lista {
		pila.Push(numero)
	}
	maximo, _ := pila.Pop()
	for !pila.IsEmpty() {
		numero, _ := pila.Pop()
		if numero > maximo {
			maximo = numero
		}
	}
	return maximo, nil
}