package main

import "fmt"

// 接口，定义一组方法
type Phone interface {
    call() int
}

type NokiaPhone struct {

}

type IPhone struct {

}

type MiPhone struct {

}

// 隐式实现接口
func (nokiaPhone NokiaPhone) call() int {
    fmt.Println("call nokia")
    return 1
}

// 隐式实现接口
func (iphone IPhone) call() int {
    fmt.Println("call iphone")
    return 2
}

// 为指针实现接口
func (miPhone *MiPhone) call() int {
    fmt.Println("call mi by pointer")
    return 3
}

// 接口值可以用作函数的参数或返回值，接口值可以看作为一个元组(value, type)
func describe(phone Phone) {
    fmt.Printf("(%v, %T)\n", phone, phone)
}

func describe2(i interface{})  {
    fmt.Printf("(%v, %T)\n", i, i)
}

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
    s := ""
    for index, ip := range ipAddr {
        fmt.Println(index, ip)
        s += string(ip)
        if index != len(ipAddr) - 1 {
            s += "."
        }
    }
    return s
}

func main()  {
    var phone Phone

    phone = new(NokiaPhone)
    describe(phone)
    phone.call()

    phone = new(IPhone)
    describe(phone)
    phone.call()

    // 定义空接口 空接口可以保存任何类型值
    var i interface{}
    describe2(i)

    i = 42
    describe2(i)

    i = "hello"
    describe2(i)

    // 判断空接口的类型
    s, ok := i.(string)
    fmt.Println(s, ok)

    f, ok := i.(float64)
    fmt.Println(f, ok)

    switch i.(type) {
    case int:
        fmt.Println("int")
    case string:
        fmt.Println("string")
    case float64:
        fmt.Println("float64")

    }

    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
