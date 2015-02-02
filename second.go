/*
 * Name: Anmol Maini
 * Date: 31st January, 2015
 * Purpose: Find the sum of the even terms of the fibionacci series under
 * 4000000.
 * Result: 4613732
 */
package main

import "fmt"

func main() {
  a, b, sum := 1, 0, 0
  for a < 4000000 {
    a, b = a+b, a
    if a%2 == 0{
      sum += a
    }
  }
  fmt.Println(sum)
}
