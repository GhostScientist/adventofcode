import Foundation

// Day 1 A Plan
// Begin by maintaining two lists, one for left numbers and one for right numbers
// For each line, parse the left and right numbers, and add to ther respective list.
// Once all lines are parsed, sort both the left and right number lists.

// Assert that there the count of left numbers must be equal to the count of right numbers
// If not, cancel the program.
// If there are, proceed.

// Maintain a totalDistance variable, and set it to 0.
// For each index in the left number list, add the absolute difference between the left and right number at that index to the totalDistance.
// Print the totalDistance.


func ingestInput() {
    while let line = readLine() {
        let split = line.split(separator: " ")
        leftNumbers.append(Int(split[0])!)
        rightNumbers.append(Int(split[1])!)
    }
}

func sortNumbers() {
    leftNumbers.sort()
    rightNumbers.sort()
}

func calculateTotalDistance() -> Int {
    var totalDistance = 0
    for i in 0..<leftNumbers.count {
        totalDistance += abs(leftNumbers[i] - rightNumbers[i])
    }
    return totalDistance
}

var leftNumbers = [Int]()
var rightNumbers = [Int]()
var totalDistance = 0 

ingestInput()
sortNumbers()
totalDistance = calculateTotalDistance()
print(totalDistance)

