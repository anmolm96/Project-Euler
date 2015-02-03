/*
 * Name: Anmol Maini
 * Date: 2nd February, 2015
 * Purpose: Find 10001st prime number.
 * Result: 104743
 */
 package main

 import  (
   "fmt"
   "math"
 )

 func isPrime(x int) bool{
   if x < 2 {
     return false
   }
   if x == 2 {
     return true
   }
   if x%2 == 0 && x>2 {
     return false
   }
   for i := 3; i < int(math.Sqrt(float64(x)) + 1); i++ {
     if x%i == 0 {
       return false
     }
   }
   return true
 }

 func main() {
   i, num := 0, 0

   for a:= 0; i < 10001; a++ {
     if isPrime(a) {
       i++
       num = a
       //fmt.Println(i, a)
     }
   }
   fmt.Println(num)
 }
