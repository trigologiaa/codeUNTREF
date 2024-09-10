package parcial1MIVI

//Dado un diccionario cuyas claves son nombres de alumnos y como valor una lista con 4 notas, Parcial 1, Recuperatorio 1, Parcial 2, Recuperatorio 2. Escribir una función que tome el diccionario y devuelva la lista de los alumnos que regularizaron la materia (aprobaron el parcial o recuperatorio 1 y el parcial o recuperatorio 2. Ej "Juan" : [4, 10, 0, 2] No aprobó, mientras que "Ana": [2, 7, 7, 0] Si aprobó

import (
	d "github.com/trigologiaa/data-structures/dictionary"
)

func Regularizacion(diccionario *d.Dictionary[string, []int]) []string {
	regularizados := []string{}
	for _, alumno := range diccionario.Keys() {
		notas, _ := diccionario.Get(alumno)
		if (notas[0] >= 4 || notas[1] >= 4) && (notas[2] >= 4 || notas[3] >= 4) {
			regularizados = append(regularizados, alumno)
		}
	}
	return regularizados
}