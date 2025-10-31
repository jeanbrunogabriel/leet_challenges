func getSneakyNumbers(nums []int) []int {
    sneakyNums :=[]int{}
    usedNums := []int{}
    for _, num := range(nums) {
        if slices.Contains(usedNums, num) && ! slices.Contains(sneakyNums, num) {
            sneakyNums = append(sneakyNums, num)
        }
        usedNums = append(usedNums, num)
    }
    return sneakyNums
}
