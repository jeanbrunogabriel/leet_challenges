func smallestNumber(n int) int {
    // 2^x - 1
    index := 0.0
    for int(math.Pow(2, index)) - 1 < n {
        index++    
    }
    nextNumber := int(math.Pow(2, index)) - 1
    return nextNumber
}
