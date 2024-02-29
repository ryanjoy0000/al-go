package main

import "math"

func binarySearchSortedArr(arr []int, val int) bool{
    result := false

    low := 0
    high := len(arr)
   
    for low < high{
        // get mid      
        mid := int(math.Round(float64(low + ((high - low) / 2))))
        // check if val found in arr
        if val == arr[mid]{
            result = true
            break
        }else if val > arr[mid]{
            // go right
            low = mid + 1
        }else{
            // go left
            high = mid
        }
    }
    return result 
}
