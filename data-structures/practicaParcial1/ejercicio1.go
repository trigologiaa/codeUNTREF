package practicaParcial1

//Implemente una función que reciba una cadena como parámetro y verifique si los paréntesis, corchetes y llaves están equilibrados en la cadena. Utilice una pila para verificar esto.

import(
	s "github.com/trigologiaa/data-structures/stack"
)

func Equilibrado(cadena string) bool {
	pila := s.NewStack[rune]()
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