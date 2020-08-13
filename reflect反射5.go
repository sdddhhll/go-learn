package main
import "fmt"
import "reflect"

type Cal struct{
	Num1 int
	Num2 int
}

func(cal Cal) GetSub(name string) int{
		return cal.Num1 - cal.Num2
}

func testReflect(b interface{}){
	//typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	var pa [2]int 
	valnum := val.NumField();
	for i:=0;i<valnum;i++ {
		tmp := val.Field(i)
		vtmp := tmp.Interface()
		vtmp2 := vtmp.(int)
		pa[i]=vtmp2
	}
	
	var para []reflect.Value
	para = append(para,reflect.ValueOf("张三"))
	res := val.Method(0).Call(para)
	tmp := res[0].Interface()
	vtmp := tmp.(int)
	fmt.Println(para[0],"完成了减法运算",pa[0],"-",pa[1],"=",vtmp)

}

func main(){
	var cal Cal = Cal{Num1:8,Num2:5}
	testReflect(cal)
}


























