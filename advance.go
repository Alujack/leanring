package main
import ("fmt"
"time")

func main(){
    for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println("Goroutine:", n)
    }(i)
    }
    time.Sleep(time.Second)
    ch := make(chan int)
    go func() { ch <- 42 }()
    val := <-ch
    fmt.Println(val)


}