package practicaParcial2

//Dadas n tareas, cada una de las cuales tarda un tiempo t en ejecutarse y un procesador que ejecuta dichas tareas, se debe diseñar una función que nos devuelva una planificación de las tareas (un orden para ejecutarlas), tal que minimice el tiempo medio de finalización (promedio de los tiempos en que finalizan las tareas). La solución debe ser O(N log N). Justificar.

type Tarea struct {
	ID		int
	Tiempo	int
}

func OrdenarTareas(tareas []Tarea) []Tarea {
	if len(tareas) < 2 {
		return tareas
	}
	medio := len(tareas) / 2
	izquierda := OrdenarTareas(tareas[:medio])
	derecha := OrdenarTareas(tareas[medio:])
	return unir(izquierda, derecha)
}

func unir(izquierda []int, derecha[]Tarea) []Tarea {
	tamaño, i, j := len(izquierda) + len(derecha), 0, 0
	largoDerecha, largoIzquierda := len(derecha) - 1, len(izquierda) - 1
	arreglo := make([]int, tamaño)
	for k := 0; k < tamaño; k++ {
		if i > largoIzquierda && j <= largoDerecha {
			arreglo[k] = derecha[j]
			j++
		} else if j > largoDerecha && i <= largoIzquierda {
			arreglo[k] = izquierda[i]
			i++
		} else if izquierda[i].Tiempo <= derecha[j].Tiempo {
			arreglo[k] = izquierda[i]
			i++
		} else {
			arreglo[k] = derecha[j]
			j++
		}
	}
	return arreglo
}