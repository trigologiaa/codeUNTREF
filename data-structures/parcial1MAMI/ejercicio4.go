package parcial1MAMI

//Escribir una función recursiva, que use la técnica de división y conquista, que reciba como parámetro una cadena de caracteres y un carácter (también como string) y devuelve la cantidad de veces que aparece dicho carácter en la cadena. Calcular el orden de la misma utilizando el teorema del maestro. Comparar con el orden teórico de una solución trivial que recorre cada elemento de la cadena y va acumulando las repeticiones del carácter buscado.

func ContarOcurrencias(s string, c string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	izquierda := s[:n / 2]
	derecha := s[n / 2:]
	return ContarOcurrencias(izquierda, c) + ContarOcurrencias(derecha, c)
}

/*
Análisis de Complejidad con Teorema del Maestro
Para aplicar el teorema del maestro, identificamos los componentes clave de la recursión:

Dividir: La cadena se divide en dos partes en cada llamada recursiva, por lo tanto 
𝑎
=
2
a=2.
Resolver: Cada llamada recursiva opera sobre aproximadamente la mitad de la cadena original, por lo que 
𝑏
=
2
b=2.
Combinar: Se realiza una suma de los resultados de las dos llamadas recursivas, lo cual toma tiempo constante 
𝑂
(
1
)
O(1).
La complejidad de la función se puede expresar como 
𝑇
(
𝑛
)
=
2
⋅
𝑇
(
𝑛
/
2
)
+
𝑂
(
1
)
T(n)=2⋅T(n/2)+O(1).

Comparando con el teorema del maestro:

log
⁡
𝑏
𝑎
=
log
⁡
2
2
=
1
log 
b
​
 a=log 
2
​
 2=1.
Como 
log
⁡
𝑏
𝑎
=
1
log 
b
​
 a=1, y 
𝑓
(
𝑛
)
=
𝑂
(
1
)
f(n)=O(1).
Entonces, según el teorema del maestro (caso 2), la complejidad de la función contarOcurrencias es 
𝑂
(
𝑛
log
⁡
𝑏
𝑎
)
=
𝑂
(
𝑛
)
O(n 
log 
b
​
 a
 )=O(n).
*/
func ContarOcurrenciasTrivial(s string, c string) int {
	ocurrencias := 0
	for _, char := range s {
		if string(char) == c {
			ocurrencias++
		}
	}
	return ocurrencias
}

/*
Análisis de Complejidad de la Solución Trivial
La solución trivial recorre cada carácter de la cadena una vez y realiza una comparación para contar las ocurrencias del carácter dado. Por lo tanto, su complejidad es 
𝑂
(
𝑛
)
O(n), donde 
𝑛
n es la longitud de la cadena. Esto se debe a que realiza una operación constante 
𝑂
(
1
)
O(1) por cada uno de los 
𝑛
n caracteres de la cadena.
*/

/*
Comparación de Eficiencias
Ambas funciones tienen complejidad 
𝑂
(
𝑛
)
O(n). Sin embargo, la función recursiva (contarOcurrencias) puede requerir más espacio en la pila de llamadas debido a las llamadas recursivas, mientras que la solución trivial (contarOcurrenciasTrivial) utiliza un bucle simple y no tiene este tipo de sobrecarga.

En términos prácticos, para cadenas grandes, la solución trivial podría ser más eficiente debido a la menor carga en la memoria y el manejo directo de los datos. La función recursiva es interesante desde el punto de vista académico y muestra un enfoque de división y conquista, pero en la práctica, la solución trivial es más directa y posiblemente más eficiente en muchos casos.
*/