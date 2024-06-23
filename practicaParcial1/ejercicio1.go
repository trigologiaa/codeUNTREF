package practicaParcia1

import(
	s "github.com/trigologiaa/data-structures/stack"
)

func Equilibrado(cadena string) bool {
	pila := NewStack[rune]()
	parejas := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, caracter := range cadena {
		switch caracter {
		case '(', '[', '{':
			pila.Push(caracter)
		case ')', ']', '}':
			if pila.IsEmpty() {
				return false
			}
			tope, _ := pila.Pop()
			if tope != parejas[caracter] {
				return false
			}
		}
	}
	return pila.IsEmpty()
}