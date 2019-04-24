package main

import "fmt"

// 声明常量
// 常量不能使用 := 声明
const world = "WORLD"

// 函数外必须用var声明变量
// 函数中可用 := 初始化变量
var var1, var2 int = 1, 2


// 函数
func add(x int, y int) int {
    return x + y
}

// 多个形参类型相同，除最后一个类型，其他可省略
func add2(x, y int) int {
    return x + y
}

// 多个返回值
func swap1(x, y string) (string, string) {
    return y, x
}

// 命名返回值
func split1(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func sqrt(x float64) float64 {
    z := float64(1)
    for i := 0; i < 10; i++ {
        z -= (z * z - x) / (2 * z)
    }
    return z
}

func main() {
    fmt.Println(sqrt(2))
}
