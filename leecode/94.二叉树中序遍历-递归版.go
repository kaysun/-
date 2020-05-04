package main

/*
94. 二叉树的中序遍历
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func main() {
	
}

func inorderTraversal(root *TreeNode) []int {
    var list []int
    list = traversal(root, list)
    return list
}

func traversal(root *TreeNode, list []int) []int {
    if root != nil {
        if root.Left != nil {
            list = traversal(root.Left, list)//注意赋值
        }
        list = append(list, root.Val)
        if root.Right != nil {
            list = traversal(root.Right, list)
        }
    }
    return list
}