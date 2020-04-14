package main
import "fmt"
import "time"

func main(){
	fishing()
}

func fishing(){
	var str string ;
	fmt.Println("请输入年-月-日")
	fmt.Scanf("%s",&str)
	t,_:= time.Parse("2006-01-02",str)
	ut := t.Unix()/60/60/24
	mod := ut % 5
//	fmt.Println(mod)
	if mod >0 && mod <4 {
		fmt.Println(str,"这一天应该打渔")
	}else{
		fmt.Println(str,"这一天应该晒网")
	}

//	fmt.Println(t.Unix())
}
