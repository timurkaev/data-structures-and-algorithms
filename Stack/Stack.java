package Stack;

public class Stack {
    private int maxSize;  // stack size
    private long[] stackArray;
    private int top; // top of the stack

    // constructor
    public Stack(int s) {
        maxSize = s;                    //  Determining the stack size
        stackArray = new long[maxSize]; // Creating an array
        top = -1;                       // No elements by default.
    }

    // Placing an element on top of the stack
    public void push(long j) {
        stackArray[++top] = j; // increase top and insert element
    }

    // Popping an array from the top of the stack
    public long pop() {
        return stackArray[top--]; // Removing an element and reducing top
    }

    // Reading an element from the stack top
    public long peek() {
        return stackArray[top];
    }

    public boolean isEmpty() {
        return (top == -1); // true if the stack is empty
    }

    public boolean isFull() {
        return (top == maxSize - 1); // true if the stack is full
    }
}

class StackApp {
    public static void main(String[] args) {
        Stack theStack = new Stack(10); // Create a new stack
        theStack.push(20);              // pushing elements on the stack
        theStack.push(30);
        theStack.push(40);
        theStack.push(50);

        // Until the stack is empty
        while(!theStack.isEmpty()) {
            long value = theStack.pop(); // remove elements from the stack
            System.out.println(value);
        }
    }
}