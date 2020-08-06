package main

import	"fmt"
func main(){
	arr := fbn(10)
	fmt.Println("--->",arr)	
}	

func fbn(n int) []int{
	var arr []int = make([]int,n)
	arr[0]=1
	arr[1]=1
	for i:=2;i<n;i++{
		arr[i]=arr[i-1]+arr[i-2]
	}
	fmt.Println(arr)
	return arr
}
