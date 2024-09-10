# Guía RadixSort + BucketSort

En las siguientes carpetas se encuentran las implementaciones del código suministrado en las clases:

- `/radixsort-old`, implementación simple de Radix Sort
- `/bucketsort`, implementación de ordenamiento por urnas
- `/radixsort`, implementación de Radix Sort

## Ejercicios

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para las siguientes consignas.

1. Reescribir el algortimo de BucketSort, sabiendo que los números que deberá ordenar se encuentran entre el 0 y el 9 inclusive, por lo tanto en lugar de guardar los números los cuenta y después genera el arreglo final con los elementos ordenados. Por ejemplo, si el arreglo a ordenar es: ```2, 1, 5, 7, 2, 5, 2, 8, 9 ,0 ,0, 1, 3, 5, 6.``` Los buckets o urnas deberían contener la siguiente información:

```
0 --> 2
1 --> 2
2 --> 3
3 --> 1
4 --> 0
5 --> 3
6 --> 1
7 --> 1
8 --> 1
9 --> 1
```

Tal que la salida generada sea ```0, 0, 1, 1, 2, 2, 2, 3, 5 ,5 ,5, 6, 7, 8, 9```

2. Modificar el código de RadixSort, para que en lugar de ordenar los números de menor a mayor los ordene de mayor a menor.
   
3. Escribir una función que ordene de manera eficiente un listado de código IATA de aeropuertos. Los códigos IATA son códigos de tres carácteres en mayúsculas. Por ejemplo: EZE Ezeiza, AEP Aeroparque, BRC Bariloche, GOA Genova, IPC Isla de Pascuas. Ver [Wikipedia][def]

[def]: https://es.wikipedia.org/wiki/Anexo:Aeropuertos_seg%C3%BAn_el_c%C3%B3digo_IATA

4. Realizar en lápiz y papel. Ordenar por RadixSort los siguientes números hexadecimales: ```143, F1A, A32, 0, 342, B33, DDA, C1A, E4D, F13, FFF```. Mostrar paso a paso
 