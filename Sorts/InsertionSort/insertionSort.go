package main

import "fmt"

func insertionSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {
		temp := arr[i]
		j := i
		for j > 0 && arr[j-1] >= temp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
	return arr
}

func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    sortedArr := insertionSort(arr)
    fmt.Println(sortedArr)
}