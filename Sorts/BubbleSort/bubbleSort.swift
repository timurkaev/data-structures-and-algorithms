func bubbleSort(_ arr: inout [Int]) {
    let len = arr.count
    for _ in 0..<len {
        for j in 0..<len - 1 {
            if arr[j] > arr[j + 1] {
                let temp = arr[j]
                arr[j] = arr[j + 1]
                arr[j + 1] = temp
            }
        }
    }
}

var arr = [64, 34, 25, 12, 22, 11, 90]
bubbleSort(&arr)
print("Sorted array: \(arr)")