func insertionSort(_ arr: [Int]) -> [Int] {
    var sortedArr: [Int] = arr
    let len: Int = sortedArr.count
    for i in 0..<len {
        let temp = sortedArr[i]
        var j = i
        while j > 0 && sortedArr[j - 1] >= temp {
            sortedArr[j] = sortedArr[j - 1]
            j -= 1
        }
        sortedArr[j] = temp
    }
    return sortedArr
}

var arr = [64, 34, 25, 12, 22, 11, 90]
let sortedArr = insertionSort(arr)
print(sortedArr)