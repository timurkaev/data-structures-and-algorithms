package main

import "fmt"

func selectionSort(arr []int) []int {
    len := len(arr)
    for i := 0; i < len; i++ {
        minIndex := i
        for j := i + 1; j < len; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        if minIndex > i {
            temp := arr[i]
            arr[i] = arr[minIndex]
            arr[minIndex] = temp
        }
    }

    return arr
}

func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    sortedArr := selectionSort(arr)
    fmt.Println(sortedArr)
}