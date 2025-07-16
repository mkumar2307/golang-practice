package main

import(
	"fmt"
)

func main() {
	var numbers []int
	var input int

	fmt.Println("Enter up to 10 integers (Press Enter after each, type a non-integer to stop):")
	// Read upto 10 integers from user
	for len(numbers) < 10 {
		_, err := fmt.Scan(&input)
		if err != nil {
			break //stop reading if number is not an integer
		}
		numbers = append(numbers, input)
	}

	Bubblesort(numbers)

	fmt.Println("Sorted numbers:")
	for _, num := range numbers{
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// Bubblesort function sorts a slice of integers in ascending order
func Bubblesort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

// Swap swaps the element at index i and i+1 in the slice
func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}