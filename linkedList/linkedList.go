package main

import "fmt"

type ListNode struct {
    data int
    next *ListNode
}

func CreateLinkedList(array []int) *ListNode {
    // O(n)
    if len(array) == 0 {
        return nil
    }
    head := &ListNode{array[0], nil}
    current := head
    for _, data := range array[1:] {
        current.next = &ListNode{data, nil}
        current = current.next
    }
    return head
}

func (head *ListNode) Append(value int) *ListNode {
    // O(n)
    if head == nil {
        return &ListNode{value, nil}
    }
    current := head
    for current.next != nil {
        current = current.next
    }
    current.next = &ListNode{value, nil}
    return head
}

func (head *ListNode) Prepend(value int) *ListNode {
    // O(1)
    return &ListNode{value, head}
}

func (head *ListNode) Delete(value int) *ListNode {
    // O(n)
    if head == nil {
        return head
    }
    if head.data == value {
        return head.next
    }
    previous := head
    current := head.next
    for current != nil {
        if current.data == value {
            previous.next = current.next
            current.next = nil
            break
        }
        previous = current
        current = current.next
    }
    return head
}

func (head *ListNode) Print() {
    // O(n)
    for head != nil {
        fmt.Print(head.data, " -> ")
        head = head.next
    }
    fmt.Print("nil\n")
}
