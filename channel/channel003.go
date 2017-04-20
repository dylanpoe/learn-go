package main
import (
    "fmt"
    "sync"
    "time"
)
func main() {
    var wg sync.WaitGroup
    exit := make(chan struct{})
    ch := make(chan int, 1)
    send := func(n int) {
        defer wg.Done()
        for {
            select {
            case ch <- n:
                time.Sleep(time.Second)
            case <-exit:
                return
            }
        }
    }
    wg.Add(10)
    // 10¸ögoroutineÐ´
    for i := 0; i < 10; i++ {
        go send(i)      
    }
    // Ò»¸ögoroutine¶Á
    go func() {
        for {
            n, ok := <-ch
            if !ok {
                return
            }
            fmt.Println(n)
        }
    }()
    time.Sleep(time.Second*10)
    close(exit)
    wg.Wait()
    close(ch)
}