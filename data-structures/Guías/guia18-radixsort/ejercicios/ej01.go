package ejercicios

//Reescribir el algortimo de BucketSort, sabiendo que los números que deberá ordenar se encuentran entre el 0 y el 9 inclusive, por lo tanto en lugar de guardar los números los cuenta y después genera el arreglo final con los elementos ordenados.
func Ej1(arr []int) {
	if len(arr) <= 1 {
		return
	}
	frequencies := make([]int, 10)
	for _, num := range arr {
		frequencies[num]++
	}
	var sortedArr []int
	for num := 0; num < 10; num++ {
		for frequencies[num] > 0 {
			sortedArr = append(sortedArr, num)
			frequencies[num]--
		}
	}
	copy(arr, sortedArr)
}