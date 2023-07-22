package Queue;

class Queue {
    private int maxSize; // The maximum queue size specified when the class was instantiated.
    private long[] queArray; // An array used to store the elements of the queue.
    private int front; // The index of the first element in the queue.
    private int rear; // The index of the last element in the queue.
    private int nItems; // Number of elements in the queue

    // Constructor
    public Queue(int s) {
        maxSize = s;
        queArray = new long[maxSize];
        front = 0;
        rear = -1;
        nItems = 0;
    }

    // Insert new element in Queue
    public void insert(long j) {
        if (rear == maxSize -1) {
            // Cyclic transfer
            rear = -1;
        }
        queArray[++rear] = j; // rear increment and insertion
        nItems++; // Increasing the number of elements
    }

    // Removing an element from the queue
    public long remove() {
        long temp = queArray[front++]; // Inserting and increasing the front
        if (front == maxSize) {
            // Cyclic transfer
            front = 0;
        }
            nItems--; // Reducing the number of elements
            return temp;
    }

    // Reading an item at the beginning of the queue
    public long peekFront() {
        return queArray[front];
    }

    // true if the queue is empty
    public boolean isEmpty() {
        return (nItems == 0);
    }

    // true if the queue is full
    public boolean isFull() {
        return nItems == maxSize - 1;
    }

    public int size() { // Number of items in the queue
        return nItems;
    }
}

class QueueApp {
    public static void main(String[] args) {
        Queue theQueue = new Queue(5); // queue of 5 cells

        // Inserting 4 elements
        theQueue.insert(10);
        theQueue.insert(20);
        theQueue.insert(30);
        theQueue.insert(40);

        // Extracting 4 elements
        theQueue.remove();
        theQueue.remove();
        theQueue.remove();

        // Inserting 4 more elements with cyclic transfer
        theQueue.insert(50);
        theQueue.insert(60);
        theQueue.insert(70);
        theQueue.insert(80);


        // Extraction and output of all elements
        while(!theQueue.isEmpty()) {
            long n = theQueue.remove();
            System.out.print(n);
            System.out.print(" ");
        }
        System.out.println("");
    }
}
