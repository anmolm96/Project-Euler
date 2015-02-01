package main

import "fmt"

func main() {
  a := 1
  b := 0
  sum := 0
  for a < 4000000 {
    a, b = a+b, a
    if a%2 == 0{
      sum += a
    }
  }
  fmt.Println(sum)
}
