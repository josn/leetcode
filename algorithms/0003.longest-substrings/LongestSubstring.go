package algs0003

func lengthOfLongestSubstring(s string) int {
    ans := 0
    length := len(s)

    start := 0
    dict := map[byte]int{}

    for i :=0; i < length; i++ {
        char := s[i]
        if idx, ok := dict[char]; ok {
            if idx >= start {
                start = idx + 1
            }
        }
        dict[char] = i

        curAns := i - start + 1
        if curAns > ans {
            ans = curAns
        }
    }

    return ans
}
