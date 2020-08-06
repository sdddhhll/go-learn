package main

import	"fmt"
func main(){
	var arr [5]int=[5]int{1,5,4,6,3}
	res:=bunSort(arr)
	fmt.Println(res)
}	

func bunSort(arr [5]int) [5]int{
		count:= len(arr)-1;
		for j:=0;j<count;j++ {
		for i:=0;i<count-j;i++ {
			if arr[i]>arr[i+1] {
				t:= arr[i]
				arr[i]=arr[i+1]
				arr[i+1]=t
			}
		}
		}
		return arr
}
