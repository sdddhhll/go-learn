package main
import "fmt"
import "math/rand"
import "time"
import "strconv"


func main(){
	fmt.Println(`
	@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	@                                                                  @
	@  这是一个猜数的游戏，系统会随机生成一个[1,100]之间的整数；       @
	@  你可以按照系统提示输入你猜想的数字，一次游戏你仅有10次机会哦    @
	@                                                                  @
	@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	`)
	game()
}

func game(){
	rand.Seed(time.Now().UnixNano())
	n:=rand.Intn(100)+1
	var m int = 0
	var str string 
	fmt.Print("游戏开始了，")
  label1:fmt.Print("请输入你的猜想：")
	fmt.Scanf("%s",&str);
	b,s:= isNum(str)
	if !b {
		fmt.Println("输入的必须是一个数字!")
		goto label1
	}

	for m<10{
		if s == n{
			break
		}else if s>n{
			fmt.Print("大了,")
		}else{
			fmt.Print("小了,")
		}
		m++
label2:fmt.Println("请再次输入你的猜想：")
	fmt.Scanf("%s",&str);
	b,s = isNum(str)
	if !b {
		fmt.Println("输入的必须是一个数字!")
		goto label2
	}
	}
	
		switch {
			case m == 1:
				fmt.Println("恭喜，猜对了，你真是个天才")
			case m>1 && m<=3 :
				fmt.Println("恭喜，猜对了，你真聪明，赶上我了")
			case m>=4 && m<=9:
				fmt.Println("猜对了，智商一般般啦")
			case m>9:
				fmt.Println("次数用尽了，说你点啥好呢,正确结果是",n)
		}

			  fmt.Println("\n马上开始新的一局，加油!\n")
		  	game()
}

func isNum(str string) (bool , int){
	n,err:=strconv.Atoi(str)
	if err == nil {
		return true,n
	}else{
		return false,0
	}
}
