class Queue {
    constructor(s) {
        this.maxSize = s; // The maximum queue size specified when the class was instantiated.
        this.queArray = []; // An array used to store the elements of the queue.
        this.front = 0; // The index of the first element in the queue.
        this.rear = -1; // The index of the last element in the queue.
        this.nItems = 0; // Number of elements in the queue
    }

    // Insert new element in Queue
    insert(j) {
        if (this.rear === this.maxSize - 1) {
            // Cyclic transfer
            this.rear = -1;
        }
        this.queArray[++this.rear] = j; // rear increment and insertion
        this.nItems++; // Increasing the number of elements
    }

    // Removing an element from the queue
    remove() {
        let temp = this.queArray[this.front++]; // Inserting and increasing the front
        if (this.front === this.maxSize) {
            // Cyclic transfer
            this.front = 0;
        }
        this.nItems--; // Reducing the number of elements
        return temp;
    }

    // Reading an item at the beginning of the queue
    peekFront() {
        return this.queArray[this.front];
    }

    // true if the queue is empty
    isEmpty() {
        return this.nItems === 0;
    }

    // true if the queue is full
    isFull() {
        return this.nItems === this.maxSize - 1;
    }

    // Number of items in the queue
    size() {
        return this.nItems;
    }
}

// Testing the Queue class
const theQueue = new Queue(5); // queue of 5 cells

// Inserting 4 elements
theQueue.insert(10);
theQueue.insert(20);
theQueue.insert(30);
theQueue.insert(40);

// Extracting 3 elements
theQueue.remove();
theQueue.remove();
theQueue.remove();

// Inserting 4 more elements with cyclic transfer
theQueue.insert(50);
theQueue.insert(60);
theQueue.insert(70);
theQueue.insert(80);

// Extraction and output of all elements
while (!theQueue.isEmpty()) {
    let n = theQueue.remove();
    console.log(n + " ");
}
console.log("");
