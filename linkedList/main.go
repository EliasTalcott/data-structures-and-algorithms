package main

import "fmt"

func main() {
    fmt.Println("\nCreate linked list")
    fmt.Println("------------------")
    head := CreateLinkedList([]int{1, 2, 3})
    head.Print()

    fmt.Println("\nAppend 4")
    fmt.Println("--------")
    head = head.Append(4)
    head.Print()

    fmt.Println("\nPrepend 0")
    fmt.Println("---------")
    head = head.Prepend(0)
    head.Print()

    fmt.Println("\nDelete 4 and 0")
    fmt.Println("--------------")
    head = head.Delete(4)
    head = head.Delete(0)
    head.Print()

    fmt.Println("")
}
