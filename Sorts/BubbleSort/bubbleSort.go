package main

import "fmt"

func bubbleSort(arr []int) []int {
    len := len(arr)
    for i := 0; i < len; i++ {
        for j := 0; j < len-1; j++ {
            if arr[j] > arr[j+1] {
                temp := arr[j]
                arr[j] = arr[j+1]
                arr[j+1] = temp
            }
        }
    }
    return arr
}

func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    sortedArr := bubbleSort(arr)
    fmt.Println("Sorted array:", sortedArr)
}