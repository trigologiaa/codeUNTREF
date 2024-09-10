package parcial1MAMI

//Dado un diccionario que contiene las notas de los estudiantes, escriba una función que devuelva la nota final (promedio de notas de cada alumno). Ej {"Perez" : [4,4,8,6], "Sánchez": [7,5,7,7], "Flores": [4,9,8]}. Debe devolver {"Perez": 5.50, "Sánchez": 6.50, "Flores": 7.0}

import(
	e "errors"
	d "github.com/trigologiaa/data-structures/dictionary"
)

func CalcularPromedios(notas *d.Dictionary[string, []float64]) (*d.Dictionary[string, float64], error) {
	promedio := func(numeros []float64) float64 {
		suma := 0.0
		for _, numero := range numeros {
			suma += numero
		}
		return suma / float64(len(numeros))
	}
	promedios := d.NewDictionary[string, float64]()
	for _, key := range notas.Keys() {
		notasEstudiante, err := notas.Get(key)
		if err != nil {
			return nil, e.New("error obteniendo las notas del estudiante")
		}
		promedio := promedio(notasEstudiante)
		promedios.Put(key, promedio)
	}
	return promedios, nil
}