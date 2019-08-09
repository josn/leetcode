package algs0008

import (
    "math"
)

const zero = rune('0')
const space = rune(' ')
const minus = rune('-')
const plus = rune('+')

func isNumber(c rune) bool {
    delta := c - zero
    return delta >= 0 && delta <=9
}

func myAtoi(str string) int {
    sign := 1
    signed := false
    numbers := []int{}
    for _, c := range str {
        if c == space && len(numbers) == 0 && !signed {
            continue
        }
        if c == minus && len(numbers) == 0 && !signed {
            sign *= -1
            signed = true
        } else if c == plus && len(numbers) == 0 && !signed {
            sign *= 1
            signed = true
        } else {
            if !isNumber(c) {
                break
            }
        }


        if isNumber(c) {
            numbers = append(numbers, int(c - zero))
        }
    }

    if len(numbers) == 0 {
        return 0
    }

    ans := 0

    ex := false

    for _, number := range numbers {
        ans = ans * 10 + number
        if ans > math.MaxInt32 {
            ex = true
        }
    }

    ans *= sign

    if sign == 1 && ex {
        return math.MaxInt32
    }

    if sign == -1 && ex {
        return math.MinInt32
    }

    return ans
}
