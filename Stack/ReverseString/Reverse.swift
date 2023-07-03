class Stack {
    private var maxSize: Int
    private var stackArray: [Character]
    private var top: Int

    init(_ s: Int) {
        maxSize = s
        stackArray = [Character](repeating: " ", count: maxSize)
        top = -1
    }

    func push(_ j: Character) -> Void {
        top += 1
        stackArray[top] = j
    }

    func pop() -> Character {
        let value = stackArray[top]
        top -= 1
        return value
    }

    func peek() -> Character {
        return stackArray[top]
    }

    func isEmpty() -> Bool {
        return top == -1
    }

    func isFull() -> Bool {
        return top == maxSize - 1
    }
}

class Reverser {
    private var input: String
    private var output: String

    init(_ inStr: String) {
        input = inStr
        output = ""
    }

    func doRev() -> String {
        let stackSize = input.count
        let theStack = Stack(stackSize)

        for ch in input {
            theStack.push(ch)
        }

        while !theStack.isEmpty() {
            let ch = theStack.pop()
            output.append(ch)
        }

        return output
    }
}

func getString() -> String {
    print("Enter a string ", terminator: "")
    let s = readLine() ?? ""
    return s
}

while true {
    let input = getString()
    if input.isEmpty {
        break
    }

    let theReverser = Reverser(input)
    let output = theReverser.doRev()
    print("Reverser: \(output)")
}