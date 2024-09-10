package practicaParcial1

//Se recibe un diccionario de deportistas premiados con indicación de campeonato y año del premio. Como resultado se debe devolver un diccionario cuyas claves son los años y para cada año todos los deportistas premiados y el campeonato ganado. Por ejemplo, si el diccionario que se recibe es: {“Gabriela Sabatini”:{1990: [“Individual F US Open”, "Doble Mixto Wimbledon"]], 1988: [“Dobles F US Open”], 1991: ["Dobles F Wimbledon"]}, “Steffi Graf”: {1988: [“Dobles F US Open”],1991: [“Single F Wimbledon”, "Single F Roland Garros"]}} El resultado debe ser: {1988: {“Steffi Graf”: [“Dobles F US Open”], “Gabriela Sabatini”: [“Dobles F US Open”]}, 1990: {“Gabriela Sabatini”: [“Individual F US Open”, "Doble Mixto Wimbledon"]]), 1991: {“Steffi Graf”: [“Single F Wimbledon”, "Single F Roland Garros], “Gabriela Sabatini”: [“Dobles F Wimbledon”]}}

import(
	d "github.com/trigologiaa/data-structures/dictionary"
)

func TransformarDatosPremios(entrada *d.Dictionary[string, *d.Dictionary[int, []string]]) *d.Dictionary[int, *d.Dictionary[string, []string]] {
	salida := d.NewDictionary[int, *d.Dictionary[string, []string]]()
	deportistas := entrada.Keys()
	for _, deportista := range deportistas {
		añosPremios, _ := entrada.Get(deportista)
		años := añosPremios.Keys()
		for _, año := range años {
			premios, _ := añosPremios.Get(año)
			if !salida.Contains(año) {
				salida.Put(año, d.NewDictionary[string, []string]())
			}
			diccionarioAño, _ := salida.Get(año)
			diccionarioAño.Put(deportista, premios)
			salida.Put(año, diccionarioAño)
		}
	}
	return salida
}