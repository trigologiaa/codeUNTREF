package parcial1MAJU2MA

//Implementar por división y conquista una función que reciba un arreglo de números enteros positivos desordenados y devuelva el máximo del arreglo. Aplicar el teorema del maestro y calcular el orden.

import(
	m "math"
)

func EncontrarMaximo(arreglo []int, izquierda, derecha int) int {
    if izquierda == derecha {
        return arreglo[izquierda]
    }
    medio := izquierda + (derecha-izquierda) / 2
    maximoIzquierda := EncontrarMaximo(arreglo, izquierda, medio)
    maximoDerecha := EncontrarMaximo(arreglo, medio + 1, derecha)
    return int(m.Max(float64(maximoIzquierda), float64(maximoDerecha)))
}

/*
Para calcular la complejidad temporal de la función findMax utilizando el teorema del maestro, vamos a analizar la recurrencia que describe el algoritmo de división y conquista que estamos utilizando.
La recurrencia para el algoritmo es:
T(n)=2T(n/2)+O(1)
donde:
*T(n) es el tiempo de ejecución de la función findMax para un arreglo de tamaño n.
*2T(n/2) corresponde al tiempo de ejecución de las dos llamadas recursivas a findMax para dos mitades del arreglo.
*O(1) corresponde al tiempo constante que toma la operación de comparación en cada nivel de recursión.
Para aplicar el teorema del maestro, debemos comparar esta recurrencia con la forma estándar T(n)=aT(n/b)+O(n^d):
*En nuestro caso, a=2 (número de subproblemas).
*b=2 (tamaño en el que se divide el problema original).
*d=0 (la complejidad de la operación de comparación, que es O(1)).
Comparando log b (a):
log 2 (2) = 1
⁡Como d=0, que es menor que log b (a) (que es 1), caemos en el caso 1 del teorema del maestro.
Según el teorema del maestro, para el caso 1 (f(n)=O(n^d) donde d < log b (a)):
T(n)=Θ(n^(log b (a)))=Θ(n^1)=Θ(n)
Por lo tanto, la complejidad temporal de la función findMax utilizando el enfoque de división y conquista es Θ(n), donde n es el tamaño del arreglo. Esto significa que la función encuentra el máximo elemento en un arreglo de tamaño n en tiempo lineal.
*/