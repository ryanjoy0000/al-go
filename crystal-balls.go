package main

import (
	"fmt"
	"math"
)


func solveCrystalBalls(breakingRange []bool, resultChan chan int) int{

    result := -1
    high := len(breakingRange)
    jumpDistance := int(math.Round(math.Sqrt(float64(high))))
    low := 0 
    for low < high{
        if breakingRange[low] { 
            break 
        }
        low += jumpDistance
    }

    high = low + 1
    low -= jumpDistance

    
    for low < high {
        if breakingRange[low]{ 
            result = low
            fmt.Println("solved in new way, loop repetitions: ")
            resultChan<- low
            break
        }
        low ++
    }

    return result

}

func oldWay(breakingRange []bool, resultChan chan int) int{
    result := -1

    for i:= 0; i < len(breakingRange); i++{
        if breakingRange[i]{
            result = i
            fmt.Println("solved in old way, loop repetitions:", i)
            resultChan<- i
            break
        }
    }

    return result
}
