/*
 * Name: Anmol Maini
 * Date: 31st January, 2015
 * Purpose: Find the sum of integers that are multiples of 3 and 5 under 1000.
 * Result: 233168
 */
package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 1000; i++ {
    if i%3==0 || i%5==0 {
      sum += i
    }
  }

  fmt.Println(sum)
}
