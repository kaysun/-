package main


type ListNode struct {
    Val int
    Next *ListNode
}

func main() {

}

//非递归
func reverseList(head *ListNode) *ListNode {
    //对链表为空和只有一个元素的边界条件进行判断
    if head == nil || head.Next == nil {
        return head
    }
    //注意最后一个空结点
    var pre *ListNode = nil
    cur := head
    for cur != nil {
        temp := cur.Next
        cur.Next = pre
        pre = cur
        cur = temp
    }
    return pre
}

//递归
//看到一句非常有意思的解释：
//我下一个节点后的所有节点都已经反转好了，现在就剩我和我的下一个节点 没有完成最后的反转了，所以反转一下我和我的下一个节点。
func reverseList2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    next := head.Next
    newHead := reverseList(next)
    next.Next = head
    head.Next = nil
    return newHead
}