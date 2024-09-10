package parcial1MIVI

//Escribir una función recursiva que calcule la potencia de un número entero.

func Potencia(base, exponente int) int {
	if exponente == 0{
		return 1
	}
	return base * Potencia(base, exponente - 1)
}