package Stack.MatchingParenthesesSearch;

class Stack {
    private int maxSize;
    private char[] stackArray;
    private int top;

    public Stack(int s) {
        maxSize = s;
        stackArray = new char[maxSize];
        top = -1;
    }

    public void push(char j) {
        stackArray[++top] = j;
    }

    public char pop() {
        return stackArray[top--];
    }

    public char peek() {
        return stackArray[top];
    }

    public boolean isEmpty() {
        return (top == -1);
    }
}

public class BracketChecker {
    public static boolean checkParentheses(String input) {
        Stack theStack = new Stack(input.length());

        for (int i = 0; i < input.length(); i++) {
            char ch = input.charAt(i);
            if (ch == '(' || ch == '{' || ch == '[') {
                theStack.push(ch);
            } else if (ch == ')' || ch == '}' || ch == ']') {
                if (theStack.isEmpty()) {
                    return false;
                }

                char top = theStack.pop();
                if ((ch == ')' && top != '(') || (ch == '}' && top != '{') || (ch == ']' && top != '[')) {
                    return false;
                }
            }
        }

        return theStack.isEmpty();
    }

    public static void main(String[] args) {
        String input = "{(2+3)*[5-2]}";
        boolean isBalanced = checkParentheses(input);

        if (isBalanced) {
            System.out.println("Parentheses are balanced.");
        } else {
            System.out.println("Parentheses are not balanced.");
        }
    }
}