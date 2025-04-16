package main
import "fmt"
func main() {

//Go has various value types including strings, integers, floats, booleans, etc.
//Strings, which can be added together with +.
    fmt.Println("go" + "lang")

//Integers and floats.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

//Booleans, with boolean operators.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}