package main

import "fmt"

func main(){
    fmt.Println("Popular Algorithms...")
   
    x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    val := 4

    // result := linearSearch(x, val)
    result := binarySearchSortedArr(x, val)
    fmt.Println("Search for", val, " in ", x ,": ", result)
}