package main

import "fmt"



func main() {
	arr := []int{10, 11, 12, 11, 12, 10, 10, 11}
	fmt.Println(checkBool(arr, 10, 12)) 

	arr2 := []int{10, 11, 12, 13, 14}
	fmt.Println(checkBool(arr2, 10, 12))
	
	arr3 := []int{1,2,3,2,3,1,2,1}
	fmt.Println(checkBool(arr3,1,3))

	arr_4:= []int{67,68,69,70,70,67,68,70,69,67}
	fmt.Println(checkBool(arr_4,67,70))
}



