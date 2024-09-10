package recuParcial1MAMI

//Escribir una función que devuelva el elemento máximo de una lista de enteros que utilice la técnica de división y conquista. Calcular el orden de la solución aplicando el Teorema del Maestro.

import(
	m "math"
)

func EncontrarMaximo(arreglo []int) int {
	if len(arreglo) == 0 {
		return m.MinInt64
	}
	return encontrarMaximoRecursivo(arreglo, 0, len(arreglo) - 1)
}

func encontrarMaximoRecursivo(arreglo []int, izquierda, derecha int) int {
	if izquierda == derecha {
		return arreglo[izquierda]
	}
	medio := (izquierda + derecha) / 2
	maximoIzquierda := encontrarMaximoRecursivo(arreglo, izquierda, medio)
	maximoDerecha := encontrarMaximoRecursivo(arreglo, medio + 1, derecha)
	return maximo(maximoIzquierda, maximoDerecha)
}

func maximo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Para calcular el orden de la solución utilizando el Teorema del Maestro para la función encontrarMaximoRecursivo en Go, vamos a analizar la recurrencia que describe su comportamiento.

Función de Recurrencia
La función de recurrencia para encontrarMaximoRecursivo es:

T(n)=2⋅T(n/2)+O(1)

donde:

T(n) es el tiempo de ejecución de la función encontrarMaximoRecursivo para un array de tamaño 
2⋅T(n/2) representa el tiempo de ejecución para dos llamadas recursivas en subarrays de tamaño n/2.
O(1) es el tiempo constante para combinar los resultados de las dos mitades (llamada a la función max).
Parámetros del Teorema del Maestro

Comparando la función de recurrencia con el Teorema del Maestro:

a=2: número de subproblemas en cada recursión.
b=2: factor por el cual se divide el tamaño del problema en cada recursión.
d=0: el tiempo constante O(1) para combinar los resultados.

Aplicación del Teorema del Maestro

Según el Teorema del Maestro, la complejidad de tiempo T(n) está dominada por el caso en el cual a>b. En este caso:

log b (a) = log 2 (2) = 1
⁡
Por lo tanto, según el Teorema del Maestro:

* Si f(n)=O(n^(log b(a))), donde log b (a) = 1
* f(n)=O(n^1)=O(n)
⁡
Por lo tanto, la complejidad temporal de la función encontrarMaximoRecursivo y, por ende, de la función encontrarMaximo, es O(n).

Conclusión
La función encontrarMaximo implementada utilizando la técnica de división y conquista tiene una complejidad temporal de O(n). Esto significa que puede encontrar el máximo elemento en una lista de enteros en tiempo lineal respecto al número de elementos en la lista.
*/