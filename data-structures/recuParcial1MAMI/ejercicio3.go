package recuParcial1MAMI

//Implementar un TAD denominado Glosario que permita asociar una palabra con varias definiciones. El Glosario debe proveer los métodos Agregar que permite agregarles definiciones a una palabra dada y un método Buscar que dada una palabra muestre por pantalla todas las definiciones de la misma.

import(
	f "fmt"
)

type Glosario struct {
	glosario map[string][]string
}

func NuevoGlosario() *Glosario {
	return &Glosario {
		glosario: make(map[string][]string),
	}
}

func(g *Glosario) Agregar(palabra, definicion string) {
	g.glosario[palabra] = append(g.glosario[palabra], definicion)
}

func(g *Glosario) Buscar(palabra string) {
	definiciones, encontrado := g.glosario[palabra]
	if encontrado {
		f.Printf("Definiciones de '%s':\n", palabra)
		for _, definicion := range definiciones {
			f.Printf("- %s\n", definicion)
		}
	} else {
		f.Printf("No se encontraron definiciones para la palabra '%s'.\n", palabra)
	}
}