package main

import	"fmt"
func main(){
	var arr [5]string=[...]string{"a","b","c","d","e"}
	var vname string = ""
	fmt.Print("please input ...")
	fmt.Scanf("%v",&vname)
	getByName(&arr,vname)	
}
func getByName(arr * [5]string,vname string) string{
	count:=len(*arr)
	index:=-1
	for i:=0;i<count;i++ {
		if (*arr)[i] == vname {
			index = i
		}
	}
		if index == -1 {
			fmt.Println("no")
		}else {
			fmt.Printf("get it :%v,index is %d",(*arr)[index],index)
		}
	return "OK"	
}	
