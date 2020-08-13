package main

import (
	"fmt"
)

func main(){
  A:= [7]int{10,20,30,40,50,60,70}

  i:=2
  j:=6

  var slc1[]int = A[i:j]
  fmt.Printf("1. len(slc_1):%d \n cap(slc_1):%d\n", j-i, len(A)-i)
  fmt.Println("slice==", slc1);
  fmt.Println(A)

  fmt.Println("Hello")
}