package algs00

func BinarySearch(array []int, dest int) int {
    start := 0
    stop := len(array) - 1

    for(start <= stop) {
        m := (start + stop) / 2
        item := array[m]
        if item > dest {
            stop = m - 1
        } else if item < dest {
            start = m + 1
        } else {
            return m
        }
    }

    return -1

}
