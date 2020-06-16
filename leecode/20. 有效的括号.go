package main

import (
    "strings"
)

func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    if len(s) % 2 != 0 {
        return false
    }
    var stack []uint8
    m := map[uint8]uint8{'(':')', '[':']', '{':'}'}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stack = append(stack, s[i])
        } else {
            if len(stack) == 0 {
                return false
            } else {
                c := stack[len(stack) - 1]
                if m[c] != s[i] {
                    return false
                }
                stack = stack[:len(stack) - 1]
            }
        }
    }
    return len(stack) == 0
}