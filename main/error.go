package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func sqrtWithError(x float64) (float64, error) {
    if x < 0 {
        return x, ErrNegativeSqrt(x)
    }
    z := float64(1)
    for i := 0; i < 10; i++ {
        z -= (z * z - x) / (2 * z)
    }
    return z, nil
}

func main()  {
    fmt.Println(sqrtWithError(-2))
    fmt.Println(sqrtWithError(2))
}
