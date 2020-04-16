package main
import "fmt"
func main(){
	var numArray01 [3]int=[3]int{1,2,3}
	var numArray02 =[3]int{4,5,6}
	var numArray03 [3]int=[...]int{7,8,9}
	var numArray04 [3]int=[...]int{1:7,2:8,0:9}
	var numArray05 [3]int=[3]int{1:11,2:18,0:19}
	    numArray06 :=[3]int{1:111,2:118,0:119}

	fmt.Println(numArray01)
	fmt.Println(numArray02)
	fmt.Println(numArray03)
	fmt.Println(numArray04)
	fmt.Println(numArray05)
	fmt.Println(numArray06)

	for index,value := range numArray06 {
		fmt.Println("index=",index,"value=",value)
	}

	for _,value := range numArray06 {
		fmt.Println("value=",value)
	}
}
