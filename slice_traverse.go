package main

import	"fmt"
func main(){
	var strslice  []string  = []string{"1","2","3","4","5"}
	strslice[4]="ceshi"
	fmt.Println(len(strslice))
	fmt.Println(cap(strslice))
	fmt.Println(strslice)
	
	var strslice2  []string  = make ([]string,100,100)
	strslice2[2]="ceshis44"
	fmt.Println(strslice2)

	for _,val:=range strslice {
		fmt.Println(val)
	}
}
