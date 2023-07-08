class Stack {
    constructor(s) {
        this.maxSize = s;
        this.stackArray = new Array(this.maxSize);
        this.top = -1;
    }

    push(j) {
        this.stackArray[++this.top] = j;
    }

    pop() {
        return this.stackArray[this.top--];
    }

    peek() {
        return this.stackArray[this.top];
    }

    isEmpty() {
        return this.top === -1;
    }
}

function checkParentheses(input) {
    const theStack = new Stack(input.length);
    for (let i = 0; i < input.length; i++) {
        const ch = input.charAt(i);
        if (ch === "(" || ch === "{" || ch === "[") {
            theStack.push(ch);
        } else if (ch === ")" || ch === "}" || ch === "]") {
            if (theStack.isEmpty()) {
                return false;
            }
            const top = theStack.pop();
            if (
                (ch === ")" && top !== "(") ||
                (ch === "}" && top !== "{") ||
                (ch === "]" && top !== "[")
            ) {
                return false;
            }
        }
    }
    return theStack.isEmpty();
}

const input = "{(2+3)*[5-2]}";

const isBalanced = checkParentheses(input);

if (isBalanced) {
    console.log("Parentheses are balanced.");
} else {
    console.log("Parentheses are not balanced.");
}
