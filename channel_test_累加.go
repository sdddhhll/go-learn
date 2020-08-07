package main
import "fmt"

/*

启动一个协程，将1-2000个数据放入一个channel中，比如numChan
启动8个协程，从numChan取出数据（比如n） 并计算1+2+...+n的值，并存放在rstChan中
最后8个协程协同完成工作后，遍历rstChan，显示结果，如：res[1]=1 ....
注意考虑rstChan int是否合适

*/

var (
	numChanCount int = 2000
	restChanCount int64 = 2000
	extChanCount int = 8
	sumMap map[int]int64 = make(map[int]int64,2000)
)




func main(){
		var numChan chan int = make(chan int,numChanCount)
		var rstChan chan map[int]int64 = make(chan map[int]int64,restChanCount)
		var extChan chan bool = make(chan bool,extChanCount)
		go func(numChan chan int){
				for i:=1;i<=numChanCount;i++ {
					numChan<-i
				}
				close(numChan)
		}(numChan)
		for i:=0 ;i<extChanCount;i++{
			go calcData(numChan,rstChan,extChan)
		}
		
		go checkExtChan(rstChan,extChan)
		
		
		for {
			sum,ok := <- rstChan
			
			if !ok {
				break
			}
			for key,value := range sum{
				sumMap[key] = value
			}
		}
		
		
		for key,value := range sumMap{
			fmt.Printf("map[%v]=%v\n",key,value)
		}
		
		fmt.Println("完成")
}

func  calcData (numChan chan int,rstChan chan map[int]int64,extChan chan bool){
	for {
		num ,ok := <- numChan
		if !ok {
			fmt.Println("关闭")
			extChan <- true
			break
		}else{
			sum := calcPlus(num)
			var sumMap map[int]int64 = make(map[int]int64,1)
			sumMap[num] = sum
			rstChan <- sumMap
		}
	}		
}

func calcPlus(n int) int64{
	num := int64(n)
	var sum int64 = 0
	var i int64 = 1
	for i=1;i<=num;i++{
		sum += i
	}
	
	return sum
}


func checkExtChan(rstChan chan map[int]int64,extChan chan bool){
	//阻塞 
	for i:=0;i<extChanCount ;i++{
		<- extChan
	}
	
	close(extChan)
	close(rstChan)
	
}


















