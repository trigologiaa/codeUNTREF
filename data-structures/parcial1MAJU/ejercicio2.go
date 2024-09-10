package parcial1MAJU

//Escribir una función que recibe un diccionario cuyas claves son los nombres de los alumnos y como valor una lista con los días que asistieron a clase. Debe devolver un diccionario con clave fecha y valor la lista de alumnos que asistieron en dicha fecha. Por ejemplo, si la entrada es {"Ana": [ "Mie 10", "Vie 12"], "Luz": [ "Vie 12", "Mie 17"], "Pedro": ["Mie 10", "Mie 17"]}, debe devolver {“Mie 10”: [“Ana", "Pedro"]), “Vie 12”: [“Ana", "Luz”], “Mie 17”: [“Luz”, "Pedro"]}

import(
	d "github.com/trigologiaa/data-structures/dictionary"
)

func InformacionPorFecha(entrada *d.Dictionary[string, []string]) *d.Dictionary[string, []string] {
	retorno := d.NewDictionary[string, []string]()
	for _, alumno := range entrada.Keys() {
		diasAsistencia, _ := entrada.Get(alumno)
		for _, dia := range diasAsistencia {
			if retorno.Contains(dia) {
				alumnosEnFecha, _ := retorno.Get(dia)
				alumnosEnFecha = append(alumnosEnFecha, alumno)
				retorno.Put(dia, alumnosEnFecha)
			} else {
				retorno.Put(dia, []string {
					alumno,
				})
			}
		}
	}
	return retorno
}