class Stack {
    private var maxSize: Int
    private var stackArray: [Character]
    private var top: Int = -1

    init(_ s: Int) {
        maxSize = s
        stackArray = Array(repeating: Character(" "), count: maxSize)
        top = -1
    }

    func push(_ j: Character) -> Void {
        top += 1
        stackArray[top] = j
    }

    func pop() -> Character {
        let element = stackArray[top]
        top -= 1
        return element
    }

    func peek() -> Character {
        return stackArray[top]
    }

    func isEmpty() -> Bool {
        return top == -1
    }
}

func checkParentheses(_ input: String) -> Bool {
    let theStack = Stack(input.count)

    for ch in input {
        if ch == "(" || ch == "{" || ch == "[" {
            theStack.push(ch)
        } else if ch == ")" || ch == "}" || ch == "]" {
            if theStack.isEmpty() {
                return false
            }

            let top = theStack.pop()
            if (ch == ")" && top != "(") ||
               (ch == "}" && top != "{") ||
               (ch == "]" && top != "[") {
                return false
               }
        }
    }

    return theStack.isEmpty()
}

let input = "{(2+3)*[5-2]}"
let isBalanced = checkParentheses(input)

if isBalanced {
    print("Parentheses are balanced")
} else {
    print("Parentheses are not balanced")
}