package practicaParcial1

//Implementar un TAD PilaCartas para jugar al solitario que permite apilar cartas de Póker (caracterizadas por el palo y el número). La carta solo se puede apilar si la pila está vacía o si en el tope de la pila se encuentra la carta inmediatamente superior del mismo palo. Por ejemplo la J de diamantes se puede apilar en una pila vacía o en una pila que en el tope tenga a la Q de diamantes.

import(
	e "errors"
	s "github.com/trigologiaa/data-structures/stack"
)

type Palo string

const (
	Corazones Palo = "Corazones"
	Diamantes Palo = "Diamantes"
	Treboles  Palo = "Treboles"
	Picas     Palo = "Picas"
)

type Carta struct {
	Palo  Palo
	Valor int
}

type PilaCartas struct {
	pila *s.Stack[Carta]
}

func NuevaPilaCartas() *PilaCartas {
	return &PilaCartas {
		pila: s.NewStack[Carta](),
	}
}

func (p *PilaCartas) PuedeApilar(carta Carta) bool {
	if p.pila.IsEmpty() {
		return true
	}
	tope, _ := p.pila.Top()
	return tope.Palo == carta.Palo && tope.Valor == carta.Valor + 1
}

func (p *PilaCartas) Apilar(carta Carta) error {
	if p.PuedeApilar(carta) {
		p.pila.Push(carta)
		return nil
	}
	return e.New("no se puede apilar la carta")
}

func (p *PilaCartas) Desapilar() (Carta, error) {
	return p.pila.Pop()
}

func (p *PilaCartas) VerTope() (Carta, error) {
	return p.pila.Top()
}

func (p *PilaCartas) EstaVacia() bool {
	return p.pila.IsEmpty()
}