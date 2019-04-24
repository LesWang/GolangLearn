package main

import "fmt"

func swap(x *int, y *int) {
    /*var temp int
    temp = *x
    *x = *y
    *y = temp*/
    *x, *y = *y, *x
}

func main()  {
    var emptyPtr *int
    fmt.Println("空指针的值: ", emptyPtr)

    var a int = 20
    var aPtr *int
    var pPtr **int
    aPtr = &a
    pPtr = &aPtr
    fmt.Println("变量地址: ", &a)
    fmt.Println("指针地址: ", aPtr)
    fmt.Println("指针指向变量值: ", *aPtr)
    fmt.Println("指针指向指针变量值: ", **pPtr)

    arr := []int {10, 100, 1000}
    var arrPtr [3]*int
    for i := 0;  i < 3; i++ {
        arrPtr[i] = &arr[i]
    }
    for i := 0; i < 3; i++ {
        fmt.Printf("arr[%d] = %d\n", i, *arrPtr[i])
    }

    var a1 = 10
    var a2 = 20
    fmt.Printf("a1=%d, a2=%d\n", a1, a2)
    swap(&a1, &a2)
    fmt.Printf("a1=%d, a2=%d\n", a1, a2)
}
