func reverse(x int) int {
    num := x
    reversedNum := 0
    for  num != 0 {
        lastDigit := num % 10
        reversedNum = reversedNum * 10 + lastDigit
        num /= 10
    }


    if reversedNum > int(math.Pow(2, 31)) -1 || reversedNum < int(math.Pow(2, 31)) * -1 {
        return 0
    }
    return reversedNum
}

