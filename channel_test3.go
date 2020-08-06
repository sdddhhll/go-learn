package main
import "fmt"
import "time"

var mapData map[int]int = make(map[int]int,10)

func calcData(exitChan chan bool,valueChan chan int){
	for{
	v,ok := <-valueChan
	var res int = 1
	if !ok{
		exitChan<-true
		close(exitChan)
		break
	}else{
		for i:=1;i<=v;i++{
		res =res * i
		mapData[v]=res
	}
	}
	time.Sleep(time.Second)
	}
}

func writeData(n int,valueChan chan int,end bool){
	fmt.Println("写入数据：",n)
	valueChan<-n
	if end {
		close(valueChan)
	}
}

func main(){
	var valueChan chan int = make(chan int,2)
	var exitChan chan bool = make(chan bool,1)
	go calcData(exitChan,valueChan)
	for i:=1;i<=10;i++{
	if i==10{
		writeData(i,valueChan,true)
	}else{
		writeData(i,valueChan,false)
	}
		
	}
	
	for {
		_,ok := <-exitChan
		if !ok {
			for index,val := range mapData{
				fmt.Printf("map[%d]：%d\n",index,val)
			}
			break
		}
	}
}