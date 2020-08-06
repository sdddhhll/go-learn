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
	slice:=num[1:3]
	fmt.Println(num)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
