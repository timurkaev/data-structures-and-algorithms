class Stack {
    private var maxSize: Int      // stack size
    private var stackArray: [Int]
    private var top: Int          // top of the stack

    // constructor
    init(_ s: Int) {
        maxSize = s;
        stackArray = [Int](repeating: 0, count: maxSize)
        top = -1
    }

    // Placing an element on top of the stack
    func push(_ j: Int) -> Void {
        // increase top and insert element
        top += 1
        stackArray[top] = j
    }

    // Popping an array from the top of the stack
    func pop() -> Int {
        let value: Int = stackArray[top]
        // Removing an element and reducing top
        top -= 1
        return value
    }

    // Reading an element from the stack top
    func peek() -> Int {
        return stackArray[top]
    }

    func isEmpty() -> Bool {
        return top == -1 // true if the stack is empty
    }

    func isFull() -> Bool {
        return top == maxSize - 1 // true if the stack is full
    }
}

let stackApp = {
    let theStack: Stack = Stack(10) // Create a new stack
    theStack.push(20); // pushing elements on the stack
    theStack.push(30);
    theStack.push(40);
    theStack.push(50);

    // Until the stack is empty
    while !theStack.isEmpty() {
        let value = theStack.pop() // // remove elements from the stack
        print(value)
    }
}

stackApp()