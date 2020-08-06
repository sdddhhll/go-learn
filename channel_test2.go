package main
import "fmt"
import "time"


func writeData(intChan chan int){
	for i:=0;i<50;i++ {
		intChan <- i
		fmt.Println("写入数据：",i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int,exitChan chan bool){
	for {
		a,ok := <- intChan
		fmt.Println("读取数据：",a)
		if !ok {
			fmt.Println("读取数据完成")
			exitChan <- true
			close(exitChan)
			break
		}
		time.Sleep(time.Second)
	}
}

func main(){
	var intChan chan int = make(chan int,50)
	var exitChan chan bool = make(chan bool,1)
	go writeData(intChan)
	go readData(intChan,exitChan)
	for {
		_,ok := <- exitChan
		if !ok{
			break
		}
		time.Sleep(time.Second)
	}
}