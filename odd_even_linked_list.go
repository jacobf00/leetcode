package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

// Helper function to print the linked list
func printList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Print(current.Val)
        if current.Next != nil {
            fmt.Print(" -> ")
        }
        current = current.Next
    }
    fmt.Println()
}

func main() {
	head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}
    head.Next.Next.Next.Next = &ListNode{Val: 5}
	
	fmt.Println(head.Next.Next.Next.Next.Next)

	// Print the original list
    fmt.Println("Original list:")
    printList(head)

    // Delete the middle node
    head = oddEvenList(head)

    // Print the modified list
    fmt.Println("After reordering:")
    printList(head)
}

// func oddEvenList(head *ListNode) *ListNode {
// 	current := head
// 	prev := head
// 	lastEven := head
// 	lastEvenPrev := head

// 	for current != nil {
// 		current = current.Next
// 		if prev.Val % 2 == 0 {
// 			lastEvenPrev = lastEven
// 			lastEven = prev
// 		}
// 		if current.Val % 2 != 0 && lastEven.Val % 2 == 0 {
// 			temp := lastEven
// 			prev.Next = lastEven
// 			lastEven.Next = current.Next
// 			lastEvenPrev.Next = current
// 			current.Next = temp.Next
// 			current = lastEven.Next
// 		}
// 		prev = prev.Next

// 	}
// 	return head
// }

func oddEvenList(head *ListNode) *ListNode {
    // Handle empty list or single node
    if head == nil || head.Next == nil {
        return head
    }
    
    // Initialize odd and even pointers
    odd := head
    even := head.Next
    evenHead := even // Store the head of even list to connect later
    
    // Traverse the list while there are nodes to process
    for even != nil && even.Next != nil {
        // Connect odd node to the next odd node
        odd.Next = even.Next
        odd = odd.Next
        
        // Connect even node to the next even node
        even.Next = odd.Next
        even = even.Next
    }
    
    // Connect odd list to even list
    odd.Next = evenHead
    
    return head
}