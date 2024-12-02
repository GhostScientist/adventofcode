import Foundation

class Report {
    let levels: [Int]
    
    init(levels: [Int]) {
        self.levels = levels
    }
    
    func isSequenceSafe(_ sequence: [Int]) -> Bool {
        guard sequence.count >= 2 else { return false }
        
        
        let isIncreasing = sequence[1] > sequence[0]
        
        for i in 0..<(sequence.count - 1) {
            let difference = sequence[i + 1] - sequence[i]
            
            let correctDirection = isIncreasing ? difference > 0 : difference < 0
            let validDifference = abs(difference) >= 1 && abs(difference) <= 3
            
            if !correctDirection || !validDifference {
                return false
            }
        }
        
        return true
    }
    
    func isSafeWithDampener() -> Bool {
        
        if isSequenceSafe(levels) {
            return true
        }
        
        for i in 0..<levels.count {
            
            var dampened = Array(levels)
            dampened.remove(at: i)
            
            if isSequenceSafe(dampened) {
                return true
            }
        }
        
        return false
    }
}

var inputLines: [String] = []

while let line = readLine() {
    inputLines.append(line)
}

let reports = inputLines
    .filter { !$0.isEmpty }
    .map { line -> Report in
        let levels = line.split(separator: " ")
            .compactMap { Int($0) }
        return Report(levels: levels)
    }

let safeReportsCount = reports.filter { $0.isSafeWithDampener() }.count
print("Number of safe reports: \(safeReportsCount)")