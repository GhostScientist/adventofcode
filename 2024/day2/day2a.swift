import Foundation

class Report {
    let levels: [Int]
    
    init(levels: [Int]) {
        self.levels = levels
    }
    
    func isSafe() -> Bool {
        guard levels.count >= 2 else { return false }
        
        let isIncreasing = levels[1] > levels[0]
        
        for i in 0..<(levels.count - 1) {
            let difference = levels[i + 1] - levels[i]
            
            let correctDirection = isIncreasing ? difference > 0 : difference < 0
            
            let validDifference = abs(difference) >= 1 && abs(difference) <= 3
            
            if !correctDirection || !validDifference {
                return false
            }
        }
        
        return true
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

let safeReportsCount = reports.filter { $0.isSafe() }.count
print("Number of safe reports: \(safeReportsCount)")