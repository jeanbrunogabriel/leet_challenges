
func maximumLengthSubstring(s string) int {
    max := 1
    runes := []rune(s)
    subStr := make(map[rune]int)
    left := 0
    right := 0

    subStr[runes[0]] = 1
    for right < len(s) - 1 {
        right += 1
        if _, ok := subStr[runes[right]]; ok {
            subStr[runes[right]] += 1
        } else {
            subStr[runes[right]] = 1
        }
        for subStr[runes[right]] > 2 {
            subStr[runes[left]] -= 1
            left += 1
        }
        window := right - left + 1
        if window > max {
            max = window
        }
    }
    return max
}
