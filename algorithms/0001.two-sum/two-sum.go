package algs0001

func twoSum(nums []int , target int) []int {
    ret := []int {}
    if len(nums) < 2 {
        return ret
    }
    
    dict := map[int]int{}
    for idx, num := range nums {
        delta := target - num
        if _, ok := dict[delta]; ok {
            return []int{dict[delta], idx}
        }
        dict[num] = idx
    }
    return ret
}
