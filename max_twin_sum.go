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
    head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	
	// Print the original list
    fmt.Println("Original list:")
    printList(head)

    result := pairSum(head)

	fmt.Println(result)
}

func pairSum(head *ListNode) int {
	nodeValue := make(map[int]int)
	nodeIndex := 0
	current := head
	maxTwinSum := 0
	currentTwinSum := 0

	for current != nil {
		// add node index and value to the hashmap
		nodeValue[nodeIndex] = current.Val
		nodeIndex++
		current = current.Next
	}
	fmt.Println(nodeValue)
	
	// get length of linked list
	n := nodeIndex
	fmt.Println("n:", n)

	// loop through first half of hashmap and add values, checking for new max sum along the way
	for i := 0; i < (n/2); i++ {
		// check if twin exists and add them, comparing to maxTwinSum
		fmt.Println("i value:", nodeValue[i], "i twin value:", nodeValue[n-1-i])
		currentTwinSum = nodeValue[i] + nodeValue[n-1-i]
		if currentTwinSum > maxTwinSum {
			maxTwinSum = currentTwinSum
		}
	}
	return maxTwinSum
}