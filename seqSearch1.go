package main

import	"fmt"
func main(){
	var arr [5]int = [...]int{1,2,3,6,9}
	binarySearch(arr,0,4,2)	

}

func binarySearch(arr [5]int,start int, end int, sea int){
	middle:=(start+end)/2
	if middle < start {
		fmt.Println("empty....")
	} else{
	if arr[middle]==sea {
		fmt.Printf("find it %d,index is->%d\n",sea,middle)
	}else if arr[middle]<sea {
		binarySearch(arr,middle+1,end,sea)
	}else {
		binarySearch(arr,start,middle-1,sea)
	}
	}
}
