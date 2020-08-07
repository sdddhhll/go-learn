package main
import "fmt"


func calcMulti(intChan chan int,extChan chan bool,resChan chan map[int]int){
	for{
		num,ok := <- intChan
		if !ok {
			fmt.Println("协城线程退出")
			extChan<- true
			break
		}
		var res int = 1
		for i:=1;i<=num;i++ {
			res = res * i
		}
		var tmp map[int]int = make(map[int]int,1)
		tmp[num]=res
		resChan<-tmp
	}
}

func writeData(intChan chan int){
	for i:=1;i<=20;i++{
		intChan<-i
	}
	close(intChan)
}

func main(){
	var intChan chan int = make(chan int , 5)
	var extChan chan bool = make(chan bool, 4)
	var resChan chan map[int]int = make(chan map[int]int,20)
	
	go writeData(intChan)
	
	for i:=0 ;i<4;i++ {
		go calcMulti(intChan,extChan,resChan)
	}
	
	go func(){
		for i:=0;i<4;i++{//阻塞
		<-extChan
	}
	close(extChan)
	close(resChan)
	}()
	
	for {
		v,ok := <-resChan 
			if !ok {
				break
			}else{
				for k1,v1 := range v {
					fmt.Printf("map[%d]=%d\n",k1,v1)
				}
			}
	}
	
	fmt.Println("结束")
	
	
}