package main
import "fmt"
import "time"

/**
	select 的错误使用方法
**/
func main(){
	chan1 := make(chan int,10)
	chan2 := make(chan string,5)
	
	go func(chan1 chan int){
		for i:=1;i<=10;i++ {
			chan1<-i
		}
	}(chan1)
	
	
	
	go func(chan1 chan int,chan2 chan string){
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
	}(chan1,chan2)


	go func(chan2 chan string){
		for i:=1;i<=5;i++{
			time.Sleep(time.Second)
			chan2 <- fmt.Sprintf("Hello :%d",i)
		}
	}(chan2)
}


















