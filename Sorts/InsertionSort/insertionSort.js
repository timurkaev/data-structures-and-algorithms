function insertionSort(arr) {
    const len = arr.length;
    for (let i = 0; i < len; i++) {
        const temp = arr[i];
        let j = i;
        while (j > 0 && arr[j - 1] >= temp) {
            arr[j] = arr[j - 1];
            j--;
        }
        arr[j] = temp;
    }
    return arr;
}

const array = [64, 34, 25, 12, 22, 11, 90];
insertionSort(array);
console.log("Sorted array: ", array);
