package main

import (
    "fmt"
)

func findMid(a, b, c int) int {
    if a <= b {
      if a <= c {
        if b <= c {
          return b
        }
        return c
      }
      return a
    }

    if a >= c {
      if b <= c {
        return c
      }
      return b
    }

    return a
}

func main() {
    var a, b, c int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    fmt.Println(findMid(a, b, c))
}
