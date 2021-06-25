package main 
import (
    "fmt"
    "time"
    "sync"
)

func main() {
    var currMap sync.Map 
    var wg sync.WaitGroup
    key1 := "key1"
    key2 := "key2" 
    key3 := "key3"
    wg.Add(3)
    go func() {
        time.Sleep(time.Second)
        currMap.Store(key1, 10)
        wg.Done()
    }()

    go func() {
        time.Sleep(time.Second)
        currMap.Store(key2, 20)
    }()

    go func() {
        time.Sleep(2 * time.Second)
        a, ok1 := currMap.Load(key1)
        if ok1 {
            fmt.Println("from key1", a)
        }
        b, ok2 := currMap.Load(key2)
        if ok2 {
            fmt.Println("from key2", b)
        }
        currMap.Store(key3, 32)
        wg.Done()
    }()

    go func() {
        time.Sleep(3 * time.Second)
        c, ok3 := currMap.Load(key3)
        if ok3 {
            fmt.Println("from key3", c)
        }
        wg.Done()
    }()
    wg.Wait()
}