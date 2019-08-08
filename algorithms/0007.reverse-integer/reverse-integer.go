package algs0007

import (
    "math"
)

func reverse(x int) int {
    sign := 1
    if x < 0 {
        x = -x
        sign = -1
    }

    ans := 0

    for (x > 0){
        mod := x % 10
        ans = ans * 10 + mod
        x = x / 10
    }
    ans = ans * sign
    if ans < math.MinInt32 || ans > math.MaxInt32 {
        return 0
    }

    return ans
}
