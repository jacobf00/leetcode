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
    head = reverseList(head)

    // Print the modified list
    fmt.Println("After reordering:")
    printList(head)
}

func reverseList(head *ListNode) *ListNode {
	next := head
	current := head
	var prev *ListNode // Explicitly declare prev as a pointer to ListNode
    
	// go until current is nil
	for current != nil {
		// move next to node after current
		next = next.Next
		// point current node toward the prev node
		current.Next = prev
		// move prev to current
		prev = current
		// move current to the next node
		current = next
	}
	return prev
}