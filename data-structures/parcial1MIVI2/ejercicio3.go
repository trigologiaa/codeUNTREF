package parcial1MIVI2

//Escribir una función recursiva que reciba un arreglo de números enteros y devuelva otro arreglo con los elementos elevados al cuadrado.

func ElevarAlCuadrado(arreglo []int) []int {
    if len(arreglo) == 0 {
        return []int{}
    }
    resultado := make([]int, len(arreglo))
    resultado[0] = arreglo[0] * arreglo[0]
    siguiente := ElevarAlCuadrado(arreglo[1:])
    copy(resultado[1:], siguiente)
    return resultado
}