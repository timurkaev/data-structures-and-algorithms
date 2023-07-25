package PriorityQueue;

import java.io.IOException;

class PriorityQueue {
    private int maxSize;
    private long[] queArray;
    private int nItems;

    // Constructor
    public PriorityQueue(int s) {
        maxSize = s;
        queArray = new long[maxSize];
        nItems = 0;
    }

    public void insert(long item) {
        int j;
        if (nItems == 0) {             // If queue is empty
            queArray[nItems++] = item; // Insert into cell 0
        } else {                       // Else queue is not empty
            for (j = nItems - 1; j >= 0; j--) { // Reverse iteration
                if (item > queArray[j]) {       // If the new element is greater,
                    queArray[j + 1] = queArray[j]; // Move top
                } else {    // if less
                    break;  // Shift stops
                }
            }
            queArray[j + 1] = item; // Insert element
            nItems++;
        }
    }

    public long remove() {
        return queArray[--nItems];
    }

    public long peekMin() {
        return queArray[nItems-1];
    }

    public boolean isEmpty() {
        return (nItems == 0);
    }

    public boolean isFull() {
        return nItems == maxSize;
    }
}

class PriorityQueueApp {
    public static void main(String[] args) throws IOException {
        PriorityQueue thePQ = new PriorityQueue(5);
        thePQ.insert(30);
        thePQ.insert(50);
        thePQ.insert(10);
        thePQ.insert(40);
        thePQ.insert(20);

        while (!thePQ.isEmpty()) {
            long item = thePQ.remove();
            System.out.print(item + " ");
        }
        System.out.println("");
    }
}