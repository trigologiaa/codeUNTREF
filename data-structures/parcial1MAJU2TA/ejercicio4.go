package parcial1MAJU2TA

//Dado un arreglo de enteros ordenados, cuyos elementos se pueden repetir, escribir una función que utilice división y conquista para devolver el índice de la primera aparición de un número buscado. Por ejemplo si a = [1,1,2,5,5,5,6,8] y se busca el 5 debe devolver 3. Si el elemento no se encuentra debe devolver -1. Calcular su orden

func PrimeraAparicion(arreglo []int, izquierda, derecha, x int) int {
    if izquierda > derecha {
        return -1
    }
    medio := izquierda + (derecha - izquierda) / 2
    if arreglo[medio] == x && (medio == izquierda || arreglo[medio - 1] < x) {
        return medio
    } else if arreglo[medio] >= x {
        return PrimeraAparicion(arreglo, izquierda, medio - 1, x)
    } else {
        return PrimeraAparicion(arreglo, medio+1, derecha, x)
    }
}