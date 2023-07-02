class Stack {
    // constructor
    constructor(s) {
        this.maxSize = s; // Determining the stack size
        this.stackArray = new Array(this.maxSize); // Creating an array
        this.top = -1; // No elements by default.
    }

    // Placing an element on top of the stack
    push(j) {
        this.stackArray[++this.top] = j; // increase top and insert element
    }

    // Popping an array from the top of the stack
    pop() {
        return this.stackArray[this.top--]; // Removing an element and reducing top
    }

    // Reading an element from the stack top
    peek() {
        return this.stackArray[this.top];
    }

    isEmpty() {
        return this.top === -1; // true if the stack is empty
    }

    isFull() {
        return this.top === this.maxSize - 1; // true if the stack is full
    }
}

const stackApp = () => {
    const theStack = new Stack(10); // Create a new stack
    theStack.push(20); // pushing elements on the stack
    theStack.push(30);
    theStack.push(40);
    theStack.push(50);

    // Until the stack is empty
    while (!theStack.isEmpty()) {
        const value = theStack.pop(); // remove elements from the stack
        console.log(value);
    }
};

stackApp();
