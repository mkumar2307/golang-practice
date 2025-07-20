package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
    id             int
    left, right    *Chopstick
    hostSemaphore  chan struct{}
}

func (p Philosopher) dine(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        // Request permission from host (blocks if 2 are already eating)
        p.hostSemaphore <- struct{}{}

        // Randomize chopstick pickup order
        if rand.Intn(2) == 0 {
            p.left.Lock()
            p.right.Lock()
        } else {
            p.right.Lock()
            p.left.Lock()
        }

        fmt.Printf("starting to eat %d\n", p.id)
        time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(100)))
        fmt.Printf("finishing eating %d\n", p.id)

        p.left.Unlock()
        p.right.Unlock()

        // Release permission
        <-p.hostSemaphore
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    chopsticks := make([]*Chopstick, 5)
    for i := range chopsticks {
        chopsticks[i] = new(Chopstick)
    }

    // Host semaphore with capacity 2
    hostSemaphore := make(chan struct{}, 2)

    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        p := Philosopher{
            id:            i + 1,
            left:          chopsticks[i],
            right:         chopsticks[(i+1)%5],
            hostSemaphore: hostSemaphore,
        }
        wg.Add(1)
        go p.dine(&wg)
    }

    wg.Wait()
}