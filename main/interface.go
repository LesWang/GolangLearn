package main

import "fmt"

type Phone interface {
    call() int
}

type NokiaPhone struct {

}

type IPhone struct {

}

func (nokiaPhone NokiaPhone) call() int {
    fmt.Println("call nokia")
    return 1
}

func (iphone IPhone) call() int {
    fmt.Println("call iphone")
    return 2
}

func main()  {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()
}
