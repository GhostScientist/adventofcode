import Foundation

func readInput() -> String {
    var input = ""
    while let line = readLine() {
        input += line + "\n"
    }
    return input
}

let input = readInput()

print("Hello, this is your input: \(input)");