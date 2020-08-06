package main

import (
	"fmt"
)
func main(){
	var num [5]int = [...]int{11,12,144,12,2}
	var sum int = 0;
	var avg float64 = 0;
	
	for _,value := range num{
		sum += value
	}
		avg = float64(sum) / float64(len(num))
	fmt.Println("sum=",sum)
	fmt.Println("avg=",avg)
	}

