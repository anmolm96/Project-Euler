/*
 * Name: Anmol Maini
 * Date: 2nd February, 2015
 * Purpose: Square of first 100 numbers - sum of squares of 100 numbers
 * Result: 25164150
 */
 package main

 import "fmt"

 func main() {
   sum1, sum2 := 0, 0

   for i:= 1; i < 101; i++ {
     sum1 += i
     sum2 += (i*i)
   }
   sum1 *= sum1
   fmt.Println(sum1-sum2)
 }
