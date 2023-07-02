func selectionSort(_ arr: [Int]) -> [Int] {
    var sortedArr = arr
    let len = sortedArr.count
    for i in 0..<len {
        var minIndex = i
        for j in (i + 1)..<len {
            if sortedArr[j] < sortedArr[minIndex] {
                minIndex = j
            }
        }
        if minIndex != i {
            let temp = sortedArr[i]
            sortedArr[i] = sortedArr[minIndex]
            sortedArr[minIndex] = temp
        }
    }
    return sortedArr
}

var arr = [64, 34, 25, 12, 22, 11, 90]
let sortedArr = selectionSort(arr)
print(sortedArr)