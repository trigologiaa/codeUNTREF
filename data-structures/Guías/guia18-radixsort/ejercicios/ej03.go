package ejercicios

// Escribir una función que ordene de manera eficiente un listado de código IATA de aeropuertos. Los códigos IATA son códigos de tres carácteres en mayúsculas.
func Ej03(arreglo []string) {
	radixSortStrings(arreglo)
}

func radixSortStrings(arreglo []string) {
	if len(arreglo) <= 1 {
		return
	}
	largoMaximo := 0
	for _, str := range arreglo {
		if len(str) > largoMaximo {
			largoMaximo = len(str)
		}
	}
	for i := largoMaximo - 1; i >= 0; i-- {
		countSortStrings(arreglo, i)
	}
}

func countSortStrings(arreglo []string, posicion int) {
	const tamañoBucket = 26
	buckets := make([][]string, tamañoBucket)
	for _, str := range arreglo {
		indice := getCharIndex(str, posicion)
		buckets[indice] = append(buckets[indice], str)
	}
	ind := 0
	for _, bucket := range buckets {
		for _, str := range bucket {
			arreglo[ind] = str
			ind++
		}
	}
}

func getCharIndex(str string, posicion int) int {
	if posicion < 0 || posicion >= len(str) {
		return 0
	}
	return int(str[posicion] - 'A')
}