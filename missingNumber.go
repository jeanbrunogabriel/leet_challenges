func missingNumber(nums []int) int {
    sum := 0

    for i := 0; i < len(nums); i++ {
        sum += (i + 1 - nums[i])
    }
    return sum 
}
