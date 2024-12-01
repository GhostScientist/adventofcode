import Foundation

// Day 1 B Plan
// Ingest left and right numbers
// For each number in left numbers,
// Does right contain left number
// If so, how many times? Multiply by the number of times it appears in the right list and add to similarity score  

func ingestInput() {
    while let line = readLine() {
        let split = line.split(separator: " ")
        print(split)
        leftNumbers.append(Int(split[0])!)
        rightNumbers.append(Int(split[1])!)
    }
}

var leftNumbers = [Int]()
var rightNumbers = [Int]()

var similarityScore = 0

func processInput () {
    for number in leftNumbers {
        if rightNumbers.contains(number) { 
            let rightCount = rightNumbers.filter { $0 == number }.count
            similarityScore += number * rightCount
        }
    }
}

ingestInput()
processInput()

print(similarityScore)
