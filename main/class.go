package main

import (
    "fmt"
    "math"
)

// go中没有类，只能定义结构体
type Vertex struct {
    x, y float64
}

// 可以通过带有接受者参数（ (v Vertex) ）的函数来为结构体定义方法
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.x * v.x + v.y * v.y)
}

// 值接受者的方法只能对Vertex值的副本进行操作
func (v Vertex) Scale(f float64) {
    v.x *= f
    v.y *= f
}

// 指针接收者的方法能够直接修改Vertex的值
func (v *Vertex) PtrScale(f float64) {
    v.x *= f
    v.y *= f
}

// 也可以为非结构体定义方法
type myFloat float64

func (f myFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    v := Vertex{3, 4}
    v.Scale(10)
    fmt.Println(v.Abs())
    v.PtrScale(10)
    fmt.Println(v.Abs())

    f := myFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}
