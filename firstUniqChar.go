func firstUniqChar(s string) int {
    hash := make(map[rune][]int)
    runes := []rune(s)
    uniqPosition := -1

    for id, r := range(runes) {
        if _, hasRune := hash[r]; ! hasRune {
            hash[r] = []int{id, 1}
        } else {
            hash[r][1] += 1  
        }
    }
    
    for _, v := range(hash) {
        if v[1] == 1 {
            if uniqPosition == -1 || v[0] < uniqPosition {
                uniqPosition = v[0]
            }
        }
    }
    return uniqPosition
}
