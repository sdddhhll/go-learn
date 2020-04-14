package main
import "fmt"

func main(){
	primeNum()
}

func primeNum(){
	var m int = 1
	for n:=1;n<=100;n++ {
		var flag int = 0
		for i:=2;i<n;i++ {
			if n % i == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 {
			if m % 5 == 0 {
				fmt.Printf("%d \n",n)
			}else{
				fmt.Printf("%d ",n)
			}
			m++
		}
	}
	fmt.Println()
}