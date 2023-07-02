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

    isFull() {
        return this.top === this.maxSize - 1;
    }
}

class Reverser {
    constructor(inStr) {
        this.input = inStr;
        this.output = "";
    }

    doRev() {
        const stackSize = this.input.length;
        const theStack = new Stack(stackSize);

        for (let j = 0; j < stackSize; j++) {
            const ch = this.input.charAt(j);
            theStack.push(ch);
        }

        while (!theStack.isEmpty()) {
            const ch = theStack.pop();
            this.output += ch;
        }

        return this.output;
    }
}

function getString() {
    const readline = require("readline");
    const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout,
    });

    return new Promise((resolve) => {
        rl.question("Enter a string: ", (answer) => {
            rl.close();
            resolve(answer);
        });
    });
}

async function main() {
    while (true) {
        const input = await getString();
        if (input === "") {
            break;
        }

        const theReverser = new Reverser(input);
        const output = theReverser.doRev();
        console.log("Reversed: " + output);
    }
}

main().catch((error) => console.error(error));
