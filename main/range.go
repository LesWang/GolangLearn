package main

import "fmt"

func main()  {
    nums := []int{2, 3, 4}
    sum := 0

    fmt.Println("range array")
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum: ", sum)
    for i, num := range nums {
        fmt.Printf("nums[%d] = %d\n", i, num)
    }
    fmt.Println()

    fmt.Println("range map")
    kvs := map[string] string {"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    fmt.Println()

    fmt.Println("range string")
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
