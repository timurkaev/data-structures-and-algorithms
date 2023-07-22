class Queue {
    private var maxSize: Int // The maximum queue size specified when the class was instantiated.
    private var queArray: [Int] // An array used to store the elements of the queue.
    private var front: Int // The index of the first element in the queue.
    private var rear: Int // The index of the last element in the queue.
    private var nItems: Int // Number of elements in the queue

    // Constructor
    init(_ s: Int) {
        maxSize = s
        queArray = Array(repeating: 0, count: maxSize)
        front = 0
        rear = -1
        nItems = 0
    }

    // Insert new element in Queue
    func insert(_ j: Int) {
        if rear == maxSize - 1 {
            // Cyclic transfer
            rear = -1
        }
        rear += 1 // rear increment
        queArray[rear] = j // insertion
        nItems += 1 // Increasing the number of elements
    }

    // Removing an element from the queue
    func remove() -> Int {
        let temp = queArray[front] // Inserting and increasing the front
        front += 1
        if front == maxSize {
            // Cyclic transfer
            front = 0
        }
        nItems -= 1 // Reducing the number of elements
        return temp
    }

    // Reading an item at the beginning of the queue
    func peekFront() -> Int {
        return queArray[front]
    }

    // true if the queue is empty
    func isEmpty() -> Bool {
        return nItems == 0
    }

    // true if the queue is full
    func isFull() -> Bool {
        return nItems == maxSize - 1
    }

    // Number of items in the queue
    func size() -> Int {
        return nItems
    }
}

// Testing the Queue class
var theQueue = Queue(5) // queue of 5 cells

// Inserting 4 elements
theQueue.insert(10)
theQueue.insert(20)
theQueue.insert(30)
theQueue.insert(40)

// Extracting 3 elements
theQueue.remove()
theQueue.remove()
theQueue.remove()

// Inserting 4 more elements with cyclic transfer
theQueue.insert(50)
theQueue.insert(60)
theQueue.insert(70)
theQueue.insert(80)

// Extraction and output of all elements
while !theQueue.isEmpty() {
    let n = theQueue.remove()
    print(n, terminator: " ")
}
print("")
