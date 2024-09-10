package ejercicios

import "math"

// Modificar el código de RadixSort, para que en lugar de ordenar los números de menor a mayor los ordene de mayor a menor.
func getDigit(num, pos int) int {
	return (num / int(math.Pow10(pos))) % 10
}

func RadixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	digits := int(math.Log10(float64(max))) + 1
	for i := 0; i < digits; i++ {
		buckets := make([][]int, 10)
		for j := 0; j < 10; j++ {
			buckets[j] = make([]int, 0)
		}
		for _, num := range arr {
			digit := getDigit(num, i)
			buckets[digit] = append(buckets[digit], num)
		}
		idx := 0
		for j := 9; j >= 0; j-- {
			for _, num := range buckets[j] {
				arr[idx] = num
				idx++
			}
		}
	}
}