package main

import "fmt"

func main()  {
    var nums1 []int
    fmt.Printf("len=%d cap=%d slice=%v\n", len(nums1), cap(nums1), nums1)

    nums2 := []int {1, 2, 3, 4, 5}
    fmt.Println("1-3: ", nums2[1:3])

    var nums3 []int
    nums3 = make([]int, 5, 10)
    fmt.Printf("len=%d cap=%d slice=%v\n", len(nums3), cap(nums3), nums3)

    nums4 := nums2[:]
    fmt.Printf("len=%d cap=%d slice=%v\n", len(nums4), cap(nums4), nums4)

    nums3 = append(nums3, 5)
    nums3 = append(nums3, 6, 7, 8, 9)
    fmt.Printf("len=%d cap=%d slice=%v\n", len(nums3), cap(nums3), nums3)

    nums5 := make([]int, len(nums3), cap(nums3) * 2)
    copy(nums5, nums3)
    fmt.Printf("len=%d cap=%d slice=%v\n", len(nums5), cap(nums5), nums5)

}
