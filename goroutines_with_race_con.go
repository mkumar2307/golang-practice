package main

import (
    "fmt"
    "time"
)

var counter int = 0

func increment() {
    for i := 0; i < 1000; i++ {
        counter++
    }
}

func main() {
    go increment()
    go increment()

    time.Sleep(1 * time.Second)
    fmt.Println("Final counter value:", counter)
}

/* 
Race Condition:
Both goroutines are modifying the shared variable counter at the same time.
The operation counter++ is not atomic — it involves reading the value, incrementing it, and writing it back.
If both goroutines read the same value before either writes back, one increment can overwrite the other, leading to lost updates.

Why It Happens:
The Go runtime interleaves the execution of the two goroutines.
Without synchronization (like a mutex), there's no guarantee about the order or timing of reads and writes.
As a result, the final value of counter is unpredictable and usually less than 2000, even though each goroutine runs 1000 increments
*/