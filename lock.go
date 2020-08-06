package main
import "fmt"
import "time"
import "sync"

var (
	dataMap map[int]int = make(map[int]int,10)
	lock sync.Mutex
)
func test(n int){
	var mul int = 1
	for i:=1;i<=n;i++{
		mul *=i
	}
	
	lock.Lock()
	dataMap[n] = mul
	lock.Unlock()
}


func main(){
		for i:=1;i<=20;i++{
			go test(i)
		}
		
		time.Sleep(2*time.Second)
		lock.Lock()
		for index,tmp := range dataMap{
			fmt.Printf("map[%d]=%d\n",index,tmp)
		}
		lock.Unlock()
}