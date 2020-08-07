package main
import "fmt"

var (
	premCount int = 4000 //启动4个信道
	maxInt int = 100000

)

func main(){
	var intChan chan int = make(chan int,10000)
	var premChan chan int =make(chan int,20000)
	var extChan chan bool = make(chan bool,premCount)
	go writeData(intChan)
	for i:=0;i<premCount;i++{
		go calcData(intChan,premChan,extChan)
	}
	
	go readExtChan(extChan,premChan)
	for {
		num,ok := <- premChan
		if !ok {
			break
		}else{
			fmt.Println("找到素数：",num)
		}
	}
	fmt.Println("完成")
}

func isPrem(n int) bool{
	for i:=2;i<n;i++{
		if n%i == 0{
			return false
		}
	}
	return true
}


func writeData(intChan chan int){
	for i:=1;i<=maxInt;i++{
		intChan <- i
		//fmt.Println("写入数据： ",i)
	}
	fmt.Println("关闭intChan")
	close(intChan)
}

func calcData(intChan chan int,premChan chan int,extChan chan bool){
	for {
		num,ok := <- intChan
		
		//fmt.Println("读取数据",num)
		if !ok {
			extChan<-true
			break
		}else{
		
			if isPrem(num) {
			
				premChan <- num
			}else{
				continue
			}
		}
		
	}
}

func readExtChan(extChan chan bool,premChan chan int){
	for i:=0;i<premCount;i++{//阻塞
		<-extChan
	}
	close(extChan)
	close(premChan)
	fmt.Println("关闭extChan")
	fmt.Println("关闭premChan")
}


