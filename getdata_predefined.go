package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Animal struct with food, locomotion, and noise
type Animal struct {
    food       string
    locomotion string
    noise      string
}

// Methods for Animal
func (a Animal) Eat() {
    fmt.Println(a.food)
}

func (a Animal) Move() {
    fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
    fmt.Println(a.noise)
}

func main() {
    // Predefined animals
    animals := map[string]Animal{
        "cow":   {"grass", "walk", "moo"},
        "bird":  {"worms", "fly", "peep"},
        "snake": {"mice", "slither", "hsss"},
    }

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("> ")
        scanner.Scan()
        input := scanner.Text()
        parts := strings.Fields(input)

        if len(parts) != 2 {
            fmt.Println("Please enter two words: <animal> <action>")
            continue
        }

        animalName := parts[0]
        action := parts[1]

        animal, exists := animals[animalName]
        if !exists {
            fmt.Println("Unknown animal. Try 'cow', 'bird', or 'snake'.")
            continue
        }

        switch action {
        case "eat":
            animal.Eat()
        case "move":
            animal.Move()
        case "speak":
            animal.Speak()
        default:
            fmt.Println("Unknown action. Try 'eat', 'move', or 'speak'.")
        }
    }
}