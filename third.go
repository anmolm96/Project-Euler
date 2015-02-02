/*
 * Name: Anmol Maini
 * Date: 1st Febuary, 2015
 * Purpose: Find largest prime factor of 600851475143
 * Result: 6857
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
   num, i, factor := 600851475143, 1, 0
   for i*i < num {
     if isPrime(i) {
       if num%i == 0 {
         factor = i
       }
     }
     i++
   }
   fmt.Println(factor)
 }
