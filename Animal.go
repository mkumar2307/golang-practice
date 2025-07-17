package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Animal interface
type Animal interface {
    Eat()
    Move()
    Speak()
}

// Cow type
type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

// Bird type
type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

// Snake type
type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("> ")
        scanner.Scan()
        input := scanner.Text()
        parts := strings.Fields(input)

        if len(parts) != 3 {
            fmt.Println("Invalid command format. Use: newanimal <name> <type> or query <name> <action>")
            continue
        }

        command, name, arg := parts[0], parts[1], parts[2]

        switch command {
        case "newanimal":
            switch arg {
            case "cow":
                animals[name] = Cow{}
            case "bird":
                animals[name] = Bird{}
            case "snake":
                animals[name] = Snake{}
            default:
                fmt.Println("Unknown animal type. Use: cow, bird, or snake.")
                continue
            }
            fmt.Println("Created it!")

        case "query":
            animal, exists := animals[name]
            if !exists {
                fmt.Println("Animal not found.")
                continue
            }
            switch arg {
            case "eat":
                animal.Eat()
            case "move":
                animal.Move()
            case "speak":
                animal.Speak()
            default:
                fmt.Println("Unknown query type. Use: eat, move, or speak.")
            }

        default:
            fmt.Println("Unknown command. Use: newanimal or query.")
        }
    }
}