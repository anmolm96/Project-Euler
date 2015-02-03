/*
 * Name: Anmol Maini
 * Date: 2nd February, 2015
 * Purpose: Find sum of all prime numbers less than 2 million.
 * Result: 142913828922
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
   sum:= 0
   for i:= 0; i < 2000000; i++ {
     if isPrime(i){
       sum += i
     }
   }
   fmt.Println(sum)
 }
