package main

import (
    "fmt"
)

type Node struct {
    Value int
    Next *Node
}

func main() {
    node := new(Node)
    node.Value = 1
    node2 := new(Node)
    node2.Value = 3
    node3 := new(Node)
    node3.Value = 2
    node4 := new(Node)
    node4.Value = 3
    node5 := new(Node)
    node5.Value = 1
    node.Next = node2
    node2.Next = node3
    node3.Next = node4
    node4.Next = node5

    head := removeDuplicate(node)
    for head != nil {
        fmt.Println(head.Value)
        head = head.Next
    }
}

func removeDuplicate(head *Node) *Node {
    var p, q, r *Node
    p = head

    for p != nil {
        q = p 
        for q.Next != nil {
            if q.Next.Value == p.Value {
                r = q.Next//要删除的结点
                q.Next = r.Next
            } else {
                q = q.Next
            }
        }
        p = p.Next
    }
    return head
}