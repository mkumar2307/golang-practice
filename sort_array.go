/* Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
    "sync"
)

func main() {
    // Prompt user for input
    fmt.Println("Enter a series of integers separated by spaces:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()

    // Convert input to slice of integers
    strs := strings.Fields(input)
    nums := make([]int, len(strs))
    for i, s := range strs {
        n, err := strconv.Atoi(s)
        if err != nil {
            fmt.Println("Invalid input:", s)
            return
        }
        nums[i] = n
    }

    // Partition the array into 4 parts
    partSize := (len(nums) + 3) / 4 // ensures all elements are included
    var wg sync.WaitGroup
    sortedParts := make([][]int, 4)

    for i := 0; i < 4; i++ {
        start := i * partSize
        end := start + partSize
        if end > len(nums) {
            end = len(nums)
        }
        if start >= len(nums) {
            break
        }

        wg.Add(1)
        go func(i int, sub []int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d sorting: %v\n", i+1, sub)
            sort.Ints(sub)
            sortedParts[i] = sub
        }(i, append([]int(nil), nums[start:end]...)) // copy to avoid race
    }

    wg.Wait()

    // Merge sorted subarrays
    merged := mergeSortedArrays(sortedParts)
    fmt.Println("Final sorted array:", merged)
}

// Merges multiple sorted slices into one sorted slice
func mergeSortedArrays(arrays [][]int) []int {
    result := []int{}
    for _, arr := range arrays {
        result = merge(result, arr)
    }
    return result
}

// Merges two sorted slices
func merge(a, b []int) []int {
    result := make([]int, 0, len(a)+len(b))
    i, j := 0, 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            result = append(result, a[i])
            i++
        } else {
            result = append(result, b[j])
            j++
        }
    }
    result = append(result, a[i:]...)
    result = append(result, b[j:]...)
    return result
}