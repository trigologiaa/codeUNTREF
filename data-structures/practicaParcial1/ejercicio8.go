package practicaParcial1

//Dado un diccionario cuyas claves son materias de una carrera y los valores son arreglos de cadenas de caracteres, donde se almacenan los apellidos de los alumnos que aprobaron esa materia, se pide escribir una función que tome ese diccionario como parámetro y devuelva otro diccionario, cuyas claves sean los apellidos de los alumnos y los valores un arreglo con las materias que aprobó ese alumno. Por ejemplo dado el diccionario de entrada: {"AyP1":["Perez", "Paz", "Martinez", "Godoy"], "AyP2": ["Martinez", "Blautzik", "Rojo", "Videla"], "EDD": ["Paz", Rojo", "Gallardo"]} Debe devolver el diccionario: {"Perez": ["AyP1"], "Paz":["AyP1", "EDD"], "Martinez": ["AyP1", "AyP2"], "Godoy": ["AyP1"], "Blautzik":["AyP2"], "Rojo": ["AyP2", "EDD"], "Videla":["AyP2"], "Gallardo":["EDD"]}

import(
	d "github.com/trigologiaa/data-structures/dictionary"
)

func MateriasPorAlumno(materias *d.Dictionary[string, []string]) *d.Dictionary[string, []string] {
	resultado := d.NewDictionary[string, []string]()
	materiasKeys := materias.Keys()
	for _, materia := range materiasKeys {
		alumnos, err := materias.Get(materia)
		if err != nil {
			continue
		}
		for _, alumno := range alumnos {
			if resultado.Contains(alumno) {
				alumnoMaterias, _ := resultado.Get(alumno)
				alumnoMaterias = append(alumnoMaterias, materia)
				resultado.Put(alumno, alumnoMaterias)
			} else {
				resultado.Put(alumno, []string {
					materia,
				})
			}
		}
	}
	return resultado
}