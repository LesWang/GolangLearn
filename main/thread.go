package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func sum(s []int, c chan int)  {
    sum := 0
    for _, v := range s {
        sum += v
    }
    // 将sum发送至信道c
    c <- sum
}

func testThread() {
    // go + 函数新建线程
    go say("world")
    say("hello")
}

func main()  {
    s := []int{1, 2, 6, -3, 4, 5}

    // 定义一个信道
    c := make(chan int)

    // 定义一个带缓冲的信道
    // 当缓冲区填满后，发送方会阻塞
    // 当缓冲区为空时，接收方会阻塞
    c2 := make(chan int, 2)

    go sum(s[:len(s) / 2], c)
    go sum(s[len(s) / 2:], c)

    go sum(s[:], c2)

    // 从信道中接受值
    x := <-c
    y, ok := <-c
    fmt.Println(ok)
    z := <-c2
    fmt.Println(x, y, z)
}
