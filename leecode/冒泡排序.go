package main

import "fmt"

func main() {
    slice := []int{21, 53, 12, 77, 45, 99, 15, 20, 78}
    sliceAfterSort := bubbleSort(slice)
    for _,v := range sliceAfterSort {
        fmt.Printf("%d ",v)
    }
    fmt.Println()
}

//冒泡排序
func bubbleSort(slice []int) []int {
    for i := 0; i < len(slice); i++ {
        for j:= i+1; j< len(slice); j++ {
            if slice[i] > slice[j] {
                slice[i], slice[j] = slice[j], slice[i]
            }
        }
    }
    return slice
}