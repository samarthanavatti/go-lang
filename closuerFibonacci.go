package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
   pePrev:=0
   prev:=1
   result:= 0
   return func(x int) int{
   if x>1{
    result = prev+pePrev
	pePrev = prev
	prev = result
    }
	if x==0{
	result = 0
	}
	if x==1{
	result =1
	}
	return result
   }

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
