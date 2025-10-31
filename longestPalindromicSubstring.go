func isPalindromicSubStr(s string) bool {
    left := 0
    right := len(s) - 1

    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}


func longestPalindrome(s string) string {
    var longestSubStr string
    for i, _ := range(s) {
        for j := len(s); j >= 0; j-- {
            if i > j {
                continue
            }
            subStr := s[i:j]
            if len(subStr) > len(longestSubStr) && isPalindromicSubStr(subStr) {
                longestSubStr = subStr
            }
        }
    }
    return longestSubStr
}
