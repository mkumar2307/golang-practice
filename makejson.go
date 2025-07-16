package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your name: ")
    nameInput, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading name:", err)
        return
    }
    name := strings.TrimSpace(nameInput)

    fmt.Print("Enter your address: ")
    addressInput, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading address:", err)
        return
    }
    address := strings.TrimSpace(addressInput)

    // Create a map to hold the name and address
    person := map[string]string{
        "name":    name,
        "address": address,
    }

    // Convert the map to a JSON object
    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }

    // Print the JSON object
    fmt.Println("JSON output:", string(jsonData))
}