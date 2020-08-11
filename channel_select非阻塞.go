package main
import "fmt"

/**
	select 的使用
**/
func main(){
	chan1 := make(chan int,10)
	chan2 := make(chan string,5)
	
	for i:=1;i<=10;i++ {
		chan1<-i
	}
	
	for i:=1;i<=5;i++{
		chan2 <- fmt.Sprintf("Hello :%d",i)
	}
	
	for {
		select {
			case v:=<- chan1:
				fmt.Printf("取到了数据：%d\n",v)
			case v:=<- chan2:
				fmt.Printf("取到了数据：%s\n",v)
			default:
				fmt.Println("在也取不到任何数据了")
				return
		}
	}

}


















