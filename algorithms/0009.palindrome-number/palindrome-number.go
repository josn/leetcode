package algs0009

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    y := 0
    _x := x
    for(x > 0) {
        m := x % 10
        y = y * 10 + m
        x = x / 10
    }

    return _x == y
}
