package algs0004

func maxInt(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func minInt(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    len1 := len(nums1)
    len2 := len(nums2)
    if len1 > len2 {
        len1, len2 = len2, len1
        nums1, nums2 = nums2, nums1
    }

    imin, imax, half := 0, len1, (len1 + len2 + 1) / 2
    for(imin <= imax) {
        i := (imin + imax) / 2
        j := half - i

        if i < len1 && nums2[j-1] > nums1[i] {
            imin = i + 1
        } else if i > 0 && nums1[i - 1] > nums2[j] {
            imax = i - 1
        } else {
            max_of_left := 0
            min_of_right := 0
            if i== 0 {
                max_of_left = nums2[j - 1]
            } else if j == 0 {
                max_of_left = nums1[i - 1]
            } else {
                max_of_left = maxInt(nums1[i-1], nums2[j-1])
            }

            if (len1 + len2) % 2 == 1 {
                return float64(max_of_left)
            }

            if i== len1 {
                min_of_right = nums2[j]
            } else if j == len2 {
                min_of_right = nums1[i]
            } else {
                min_of_right = minInt(nums1[i], nums2[j])
            }

            return float64(max_of_left + min_of_right) / 2.0
        }
    }

    return 0.0
}
