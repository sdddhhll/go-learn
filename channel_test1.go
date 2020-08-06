package main
import "fmt"

type Cat struct{
	Name string
	Age int
}

func main(){
	var intChan chan int = make(chan int,3)
	intChan<-3
	intChan<-4
	intChan<-5
	
	fmt.Println("---------之前-----------")
	
	for v := range intChan {
		fmt.Println(v)
	}
	close(intChan)//内置函数，用来关闭管道
	fmt.Println("---------之后-----------")
	
	for v := range intChan {
		fmt.Println(v)
	}
}