func reverseWords(s string) string {
    runes := []rune(s)
    wordStart := 0

    for i := 0; i <= len(runes); i++ {
        if i == len(runes) || runes[i] == ' ' {
            left := wordStart
            right := i - 1
            for left < right {
                runes[left], runes[right] = runes[right], runes[left]
                left++
                right--
            }
            wordStart = i + 1
        }
    }
    return string(runes)
}
