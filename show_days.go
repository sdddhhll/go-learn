package main
import "fmt"

func main(){
	for{
		var y int;
		var m int;
		fmt.Print("请输入年份：")
		fmt.Scanf("%d",&y);
		fmt.Print("请输入月份：")
		fmt.Scanf("%d",&m)
		showDays(y,m)
	}
}

func showDays(y,m int){
	if m>12 || m<1 {
		fmt.Println("月份输入错误，请重新输入。")
		showDays(y,m)
	}else{
		switch m {
		case 1,3,5,7,8,10,12:
			fmt.Println(m,"月份共[31]天")
		case 4,6,9,11:
			fmt.Println(m,"月份共[30]天")
		case 2:
			n:=show2Days(y)
			fmt.Println(m,"月份共[",n,"]天")
		}
	}
}


func show2Days(y int) int{
	if checkYear(y) {
		return 29
	}else
	{
		return 28
	}
}

func checkYear(y int) bool{
	if y%4 != 0 {
		return false
	}else if y%4==0 && y%100!=0{
		return false
	}else{
		return true
	}
}