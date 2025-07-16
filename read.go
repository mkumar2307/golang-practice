package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Define the name struct
type Name struct {
    Fname string
    Lname string
}

func main() {
    // Prompt the user for the filename
    fmt.Print("Enter the name of the text file: ")
    var filename string
    fmt.Scanln(&filename)

    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a slice to hold Name structs
    var names []Name

    // Read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.SplitN(line, " ", 2)
        if len(parts) == 2 {
            // Trim to ensure no extra spaces and limit to 20 characters
            fname := truncateString(strings.TrimSpace(parts[0]), 20)
            lname := truncateString(strings.TrimSpace(parts[1]), 20)
            names = append(names, Name{Fname: fname, Lname: lname})
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Print the names
    fmt.Println("\nNames found in the file:")
    for _, name := range names {
        fmt.Printf("First Name: %s, Last Name: %s\n", name.Fname, name.Lname)
    }
}

// Helper function to truncate a string to a maximum length
func truncateString(s string, maxLen int) string {
    if len(s) > maxLen {
        return s[:maxLen]
    }
    return s
}