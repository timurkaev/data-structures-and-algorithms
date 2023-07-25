import Foundation

class PriorityQueue {
    private var maxSize: Int
    private var queArray: [Int64]
    private var nItems: Int

    // Constructor
    init(s: Int) {
        maxSize = s
        queArray = Array(repeating: 0, count: maxSize)
        nItems = 0
    }

    // Insert an element into the priority queue
    func insert(item: Int64) {
        var j: Int // Declare the variable j here
        if nItems == 0 {                // If the queue is empty
            queArray[nItems] = item     // Insert the item into cell 0
            nItems += 1
        } else {                       // Else, the queue is not empty
            j = nItems - 1 // Initialize j outside the for loop
            for j in stride(from: j, through: 0, by: -1) { // Reverse iteration
                if item > queArray[j] {       // If the new element is greater
                    queArray[j + 1] = queArray[j] // Move the current element up
                } else {    // If the new element is less or equal
                    break     // Shift stops
                }
            }
            queArray[j + 1] = item // Insert the element
            nItems += 1
        }
    }

    // Remove the minimum element from the priority queue
    func remove() -> Int64 {
        nItems -= 1
        return queArray[nItems]
    }

    // Peek the minimum element in the priority queue
    func peekMin() -> Int64 {
        return queArray[nItems - 1]
    }

    // Check if the queue is empty
    func isEmpty() -> Bool {
        return nItems == 0
    }

    // Check if the queue is full
    func isFull() -> Bool {
        return nItems == maxSize
    }
}

// Entry point of the application
func main() {
    let thePQ = PriorityQueue(s: 5)
    thePQ.insert(item: 30)
    thePQ.insert(item: 50)
    thePQ.insert(item: 10)
    thePQ.insert(item: 40)
    thePQ.insert(item: 20)

    while !thePQ.isEmpty() {
        let item = thePQ.remove()
        print(item, terminator: " ")
    }
    print("")
}

// Call the entry point
main()
