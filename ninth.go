/*
 * Name: Anmol Maini
 * Date: 2nd February, 2015
 * Purpose: Pythogorean triplet where a+b+c = 1000
 * Result: 31875000.
 */
 package main

 import "fmt"

 func main() {
   loop:
   for a:= 1; a < 1000; a++ {
     for b := 1; b < 1000; b++ {
       for c := 1; c < 1000; c++ {
         if c*c == a*a + b*b{
           if a+b+c == 1000 {
             fmt.Println(a*b*c)
             fmt.Println(a)
             fmt.Println(b)
             fmt.Println(c)
             break loop
           }
         }
       }
     }
   }
 }
