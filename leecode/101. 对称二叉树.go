package main

func main() {

}

/*
	101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。

 

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
 

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
 

进阶：

你可以运用递归和迭代两种方法解决这个问题吗？
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return helpSymmetry(root, root)
}

func helpSymmetry(t1, t2 *TreeNode) bool {
    if (t1 == nil && t2 != nil) ||  (t1 != nil && t2 == nil) {
        return false
    }

    //没有这个判断，60行就会报空指针错误
    if t1 == nil && t2 == nil {
        return true
    }

    return t1.Val == t2.Val && helpSymmetry(t1.Left, t2.Right) && helpSymmetry(t1.Right, t2.Left)
}