package main

import (
    "fmt"
    "strconv"
)

func compress(chars []byte) int {
    var write, anchor, num int
    var numBytes []byte
    for read, char := range chars {
        if read + 1 == len(chars) || char != chars[read + 1] {
            num = read - anchor + 1 
            
            chars[write] = chars[anchor]
            if num > 1 {
                numBytes = []byte(strconv.Itoa(num))
                for index := 0; index < len(numBytes); index ++ {
                    chars[write+index+1] = numBytes[index]
                }
                write = write + len(numBytes) + 1
            } else {
                write++
            }
            anchor = read + 1
        }
    }
    return write
}