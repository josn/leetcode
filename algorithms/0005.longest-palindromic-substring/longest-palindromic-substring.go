package algs0005

func longestPalindrome(s string) string {
    length := len(s)
    max := 0
    ans := ""

    if length > 0 {
        ans = string(s[0])
    }

    for i := 0; i < length; i++ {
        start, stop := -1, -1
        j, k := i-1, i+1
        for ; j >= 0 && k < length; j, k = j-1, k+1 {
            a := s[j]
            b := s[k]
            if a != b {
                break
            } else {
                start, stop = j, k
            }
        }

        if (start >= 0 && stop < length) {
            _max := stop - start + 1
            if _max > max {
                max = _max
                ans = s[start: stop+1]
            }
        }

    }

    for i := 0; i < length; i++ {
        start, stop := -1, -1
        j, k := i, i+1
        for ; j >= 0 && k < length; j, k = j-1, k+1 {
            a := s[j]
            b := s[k]
            if a != b {
                break
            } else {
                start, stop = j, k
            }
        }

        if (start >= 0 && stop < length) {
            _max := stop - start + 1
            if _max > max {
                max = _max
                ans = s[start: stop+1]
            }
        }

    }

    return ans
}
