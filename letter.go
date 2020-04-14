package main
import "fmt"

func main(){
	calc()
}

func calc(){
	//var a int = 'Z'
	//fmt.Printf("%d",a)
	for i:=0;i<26;i++ {
		fmt.Printf("%c",97+i)	
		fmt.Printf("%c ",90-i)	
	}
	fmt.Println()
}
