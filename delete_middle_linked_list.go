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

	// Print the original list
    fmt.Println("Original list:")
    printList(head)

    // Delete the middle node
    head = deleteMiddle(head)

    // Print the modified list
    fmt.Println("After deleting middle node:")
    printList(head)
}

func deleteMiddle(head *ListNode) *ListNode {
    // If list has 1 node, deleting it results in empty list
    if head == nil || head.Next == nil {
        return nil
    }
    
    // Initialize slow and fast pointers
    slow := head
    fast := head
    var prev *ListNode = nil // To keep track of node before slow
    
    // Move fast 2 steps and slow 1 step until fast reaches end
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        prev = slow
        slow = slow.Next
    }
    
    // At this point, slow is at the middle node
    // Link prev to slow's next to skip the middle node
    prev.Next = slow.Next
    
    return head
}