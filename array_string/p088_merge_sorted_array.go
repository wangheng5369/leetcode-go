package array_string

// merge merges two sorted arrays nums1 and nums2 into nums1.
// nums1 has length m+n, with the first m elements valid.
func merge(nums1 []int, m int, nums2 []int, n int) {
	arr := make([]int, 0, m)
	for i := 0; i < m; i++ {
		arr = append(arr, nums1[i])
	}

	l1, l2, idx := 0, 0, 0
	for l1 < m && l2 < n {
		if arr[l1] <= nums2[l2] {
			nums1[idx] = arr[l1]
			l1++
		} else {
			nums1[idx] = nums2[l2]
			l2++
		}
		idx++
	}
	for l1 < m {
		nums1[idx] = arr[l1]
		idx++
		l1++
	}
	for l2 < n {
		nums1[idx] = nums2[l2]
		idx++
		l2++
	}
}