package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    // Initialize a slice with length 0 and capacity 3
    numbers := make([]int, 0, 3)
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Enter an integer (or 'X' to quit): ")
        scanner.Scan()
        input := strings.TrimSpace(scanner.Text())

        if strings.EqualFold(input, "X") {
            fmt.Println("Exiting program.")
            break
        }

        num, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Invalid input. Please enter a valid integer or 'X' to quit.")
            continue
        }

        numbers = append(numbers, num)
        sort.Ints(numbers)
        fmt.Println("Sorted slice:", numbers)
    }
}