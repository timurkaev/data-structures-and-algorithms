package Sorts.InsertionSort;

public class InsertionSort {
    public static void insertionSort(int[] arr) {
        int len = arr.length;
        for (int i = 1; i < len; i++) {
            int temp = arr[i];
            int j = i;
            while(j  > 0 && arr[j - 1] >= temp) {
                arr[j] = arr[j - 1];
                j--;
            }
            arr[j] = temp;
        }
    }

    public static void main(String[] args) {
        int[] arr = {64, 34, 25, 12, 22, 11, 90};
        insertionSort(arr);
        System.out.println("Sorted array:");
        for (int num : arr) {
            System.out.print(num + " ");
        }
    }
}
