package recuParcial1MAMI

//Implementar un TAD Cola que internamente utiliza dos pilas para contener los datos de la siguiente manera: cada vez que se encola un elemento se lo guarda en la pila A, y cada vez que se desencola una elemento se lo extrae del tope de la pila B, si la pila B está vacía entonces se desapilan todos los elementos de la pila A y se los apila en B. Los métodos que debe soportar son Encolar, Desencolar, EstaVacia y Frente (para ver el primer elemento de la Cola).

import(
	s "github.com/trigologiaa/data-structures/stack"
)

type ColaConDosPilas[T any] struct {
	pilaA *s.Stack[T]
	pilaB *s.Stack[T]
}

func NuevaColaConDosPilas[T any]() *ColaConDosPilas[T] {
	return &ColaConDosPilas[T] {
		pilaA: s.NewStack[T](),
		pilaB: s.NewStack[T](),
	}
}

func(c *ColaConDosPilas[T]) Encolar(v T) {
	c.pilaA.Push(v)
}

func(c *ColaConDosPilas[T]) Desencolar() (T, error) {
	if c.pilaB.IsEmpty() {
		for !c.pilaA.IsEmpty() {
			elemento, _ := c.pilaA.Pop()
			c.pilaB.Push(elemento)
		}
	}
	return c.pilaB.Pop()
}

func(c *ColaConDosPilas[T]) EstaVacia() bool {
	return c.pilaA.IsEmpty() && c.pilaB.IsEmpty()
}

func(c *ColaConDosPilas[T]) Frente() (T, error) {
	if c.pilaB.IsEmpty() {
		for !c.pilaA.IsEmpty() {
			elemento, _ := c.pilaA.Pop()
			c.pilaB.Push(elemento)
		}
	}
	return c.pilaB.Top()
}