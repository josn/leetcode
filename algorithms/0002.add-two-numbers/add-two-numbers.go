package algs0002

type ListNode struct {
    Val int
    Next *ListNode
}

func makeListNode(ins []int) *ListNode {
    if len(ins) == 0 {
        return nil
    }
    ret := new(ListNode)
    head := ret
    head.Val = ins[0]
    for i := 1; i < len(ins); i++ {
        head.Next = new(ListNode)
        head = head.Next
        head.Val = ins[i]
    }

    return ret
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ret := new(ListNode)
    result := ret
    kept := 0
    
    for(l1 != nil && l2 != nil) {
        val := l1.Val + l2.Val + kept
        kept = val / 10
        val = val % 10
        ret.Val = val

        l1 = l1.Next
        l2 = l2.Next
        
        if(l1 != nil || l2 != nil || kept > 0 ) {
            ret.Next = new(ListNode)
            ret = ret.Next
        }
    }

    if (l2 != nil ) {
        l1 = l2
    }
    
    for (l1 != nil) {
        val := l1.Val + kept
        kept = val / 10
        val = val % 10

        ret.Val = val
        l1 = l1.Next
        if (l1 != nil || kept > 0) {
            ret.Next = new(ListNode)
            ret = ret.Next
        }
    }

    if kept > 0 {
        ret.Val = kept
    }
    
    return result
}
