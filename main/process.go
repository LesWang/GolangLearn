package main

import (
    "fmt"
    "time"
)

func switch1()  {
    t := time.Now()
    fmt.Println(t.Hour())
    // switch没有条件，相当于switch true
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}

func defer1() {
    // defer语句会将函数推迟到外层函数返回之后执行
    // 但推迟调用的函数，其参数会立即求值
    defer fmt.Println("in defer")

    // defer语句会将推迟调用的函数压入栈 按先进后出执行
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("out defer")
}

func main()  {
    defer1()
}
