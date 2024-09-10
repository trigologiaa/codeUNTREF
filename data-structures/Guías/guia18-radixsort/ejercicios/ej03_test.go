package ejercicios

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEj03(test *testing.T) {
	casos := []struct {
		entrada    	[]string
		esperado 	[]string
	}{
		{[]string{"EZE", "AEP", "BRC", "GOA", "IPC"}, []string{"AEP", "BRC", "EZE", "GOA", "IPC"}},
		{[]string{"JFK", "LAX", "ORD", "ATL", "DFW"}, []string{"ATL", "DFW", "JFK", "LAX", "ORD"}},
		{[]string{"SFO", "SEA", "DEN", "MIA", "PHX"}, []string{"DEN", "MIA", "PHX", "SEA", "SFO"}},
		{[]string{"LHR", "CDG", "FRA", "AMS", "BCN"}, []string{"AMS", "BCN", "CDG", "FRA", "LHR"}},
	}

	for _, caso := range casos {
		copiaEntrada := make([]string, len(caso.entrada))
		copy(copiaEntrada, caso.entrada)
		Ej03(copiaEntrada)
		assert.Equal(test, caso.esperado, copiaEntrada, "Para entrada %v", caso.entrada)
	}
}
