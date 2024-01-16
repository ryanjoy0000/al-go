package main

func linearSearch(arr []int, val int) bool{
    result := false
   
    for _, v := range arr {
        if v == val{
            result = true
            break
        }
    }

    return result 
}
