package main
import "fmt"

func main(){
	calc()
}

func calc(){
	fmt.Println(`
	1.加法
	2.减法
	3.乘法
	4.除法
	0.退出
	`);
	var n int;
	fmt.Print("输入选择:")
	fmt.Scanf("%d",&n)
	switch n {
		case 0:
			
		case 1:
				add()
		case 2:
				sub()
		case 3:
				mul()
		case 4:
				div()
		default:
			fmt.Println("输入有误，请重新输入")
	}
}

func add(){
	fmt.Printf("%d+%d=%d \n",10,5,10+5 )
}

func sub(){
	fmt.Printf("%d-%d=%d \n",10,5,10-5 )
}

func mul(){
	fmt.Printf("%d*%d=%d \n",10,5,10*5 )
}

func div(){
	fmt.Printf("%d/%d=%d \n",10,5,10/5 )
}
