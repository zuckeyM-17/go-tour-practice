package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    n := 0.0
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2 * z)
        if math.Abs(z-n) < 0.0000001 {
            return z
        }
        n = z
    }

    return z
}

func main() {
    x := 2.0
    fmt.Println(Sqrt(x))
    fmt.Println(math.Sqrt(x))
}