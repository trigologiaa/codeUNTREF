package practicaParcial1

//Implementar una cola que internamente use dos pilas para almacenar los datos.

import (
	e "errors"
	s "github.com/trigologiaa/data-structures/stack"
)

type DoubleStackQueue[T any] struct {
	stack1 *s.Stack[T]
	stack2 *s.Stack[T]
}

func NewDoubleStackQueue[T any]() *DoubleStackQueue[T] {
	return &DoubleStackQueue[T] {
		stack1: s.NewStack[T](),
		stack2: s.NewStack[T](),
	}
}

func (q *DoubleStackQueue[T]) Enqueue(valor T) {
	q.stack1.Push(valor)
}

func (q *DoubleStackQueue[T]) Dequeue() (T, error) {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			valor, err := q.stack1.Pop()
			if err != nil {
				return valor, err
			}
			q.stack2.Push(valor)
		}
	}
	if q.stack2.IsEmpty() {
		var zero T
		return zero, e.New("cola vacía")
	}
	return q.stack2.Pop()
}

func (q *DoubleStackQueue[T]) Front() (T, error) {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			valor, err := q.stack1.Pop()
			if err != nil {
				return valor, err
			}
			q.stack2.Push(valor)
		}
	}
	if q.stack2.IsEmpty() {
		var zero T
		return zero, e.New("cola vacía")
	}
	top, err := q.stack2.Top()
	if err != nil {
		return top, err
	}
	return top, nil
}

func (q *DoubleStackQueue[T]) IsEmpty() bool {
	return q.stack1.IsEmpty() && q.stack2.IsEmpty()
}
