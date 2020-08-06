package main

import	"fmt"
import	"math/rand"
import	"time"
func main(){
	var num [5]int
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<5;i++{
		num[i]=rand.Intn(100)
	}
	fmt.Println("pre:",num)
	length:=len(num)
	fmt.Println(length)
	for i:=0;i<length/2;i++{
		temp:=num[length-1-i]
		num[length-1-i]=num[i]
		num[i]=temp
	}
	fmt.Println("post:",num)
} 
