package main

import "math"

func binarySearchSortedArr(arr []int, val int) bool{
    result := false

    low := 0
    high := len(arr)
   
    canContinue := true
    for canContinue{
    
        // get mid      
        mid := int(math.Round(float64(low + ((high - low) / 2))))
    
        if val == arr[mid]{
            // val found in arr
            result = true
            canContinue = false
            break
        }else if val > arr[mid]{
            // go right
            low = mid + 1
        }else{
            // go left
            high = mid
        }
        
        // condition
        if low >= high {
            canContinue = false
        }
    }
    return result 
}
