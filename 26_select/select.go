package main

import (
    "fmt"
    "time"
)

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()

    tick := time.Tick(1000 * time.Millisecond)
    boom := time.After(5000 * time.Millisecond)
    for {
        select {
        case <-tick:
            fmt.Println("Tick.")
        case <-boom:
            fmt.Println("BOOM!")
            fibonacci(c, quit)
            return
        default:
            fmt.Println("    .")
            time.Sleep(100 * time.Millisecond)
        }
    }

}