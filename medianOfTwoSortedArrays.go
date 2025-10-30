func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    mergedArray := append(nums1, nums2...)
    mergedArrayLenght := len(mergedArray)
    var result float64
    sort.Ints(mergedArray)
    if mergedArrayLenght % 2 == 1 {
        median := mergedArrayLenght / 2
        result = float64(mergedArray[median])
    } else {
        median := mergedArrayLenght / 2
        result = (float64(mergedArray[median]) + float64(mergedArray[median-1]))
        result /= 2
    }
    return result
    
}
