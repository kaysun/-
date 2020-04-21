package main

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main()

type ListNode struct {
    Val int
    Next *ListNode
}

//反转前n个结点
 func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == 1 {
        return reverseN(head, n)
    }
    //前进到第m个结点，然后按照前“n”个结点的逻辑处理
    head.Next = reverseBetween(head.Next, m - 1, n - 1)
    return head
}

//记录结点n的后驱结点
var nextNode *ListNode = nil

func reverseN(head *ListNode, n int) *ListNode {
    if n == 1 {
        nextNode = head.Next
        return head
    }
    newHead := reverseN(head.Next, n - 1)
    //断环处理
    head.Next.Next = head
    head.Next = nextNode
    return newHead
}