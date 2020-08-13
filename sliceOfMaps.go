package main

import "fmt"

func main(){
	myMap1 := make(map[string]string)
	myMap2 := make(map[string]string)

	myMap1["a"]="hii"
	myMap1["b"]="hello"

	myMap2["p"]="p-hii"
	myMap2["q"]="q-hello"

	sliceOfMaps:= make([]map[string]string,2)
	sliceOfMaps[0] = myMap1
	sliceOfMaps[1] = myMap2

	fmt.Println(sliceOfMaps)
	for res := range sliceOfMaps{
		fmt.Printf("map==%v\n", sliceOfMaps[res]);
	}
}