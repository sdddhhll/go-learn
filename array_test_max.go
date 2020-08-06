package main

import (
	"fmt"
)
func main(){
	var num [5]int = [...]int{11,12,144,12,1}
	var index int = 0;
	var max int = 0;
	for i:=0;i<len(num);i++{
		if num[i] > max {
			max = num[i]
			index = i
		}
	}
	fmt.Printf("num[%d]=%d",index,max)
}
