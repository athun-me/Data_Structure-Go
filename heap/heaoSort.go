package main

import (
    "fmt"
)

func heapSort(arr []int) {
    n := len(arr)

    // Build max heap
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }

    // Extract elements from heap in descending order
    for i := n - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i, 0)
    }
}

// Heapify function to build max heap
func heapify(arr []int, n, i int) {
    largest := i
    l := 2*i + 1
    r := 2*i + 2

    if l < n && arr[l] > arr[largest] {
        largest = l
    }

    if r < n && arr[r] > arr[largest] {
        largest = r
    }

    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}

func main() {
    arr := []int{12, 11, 13, 5, 6, 7}
    heapSort(arr)
    fmt.Println(arr)
}
