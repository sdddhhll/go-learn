package main
import "fmt"
import "reflect"

type Student struct{
	Name string `json:"user_name"`
	Age int
	No string
	Score float64
}

//func (stu *Student) GetInfo(){  这样写不行 下同
func (stu Student) GetInfo(){
	fmt.Println("姓名：",stu.Name)
}


func (stu Student) SetScore(sc float64){
	stu.Score = sc
	fmt.Println("stu1：",stu)
}

func (stu Student) SetScoreAndName(sc float64,na string){
	stu.Name = na
	stu.Score = sc
	fmt.Println("stu2：",stu)
}

func testReflect(b interface{}){
	vTy := reflect.TypeOf(b)
	vVal := reflect.ValueOf(b)
	kd := vVal.Kind()
	if kd != reflect.Struct {
		fmt.Println("参数必须是结构体")
		return
	}
	
	numf := vVal.NumField()
	
	fmt.Println(numf)
	
	for i:=0 ;i<numf;i++{
		fmt.Printf("field=%v,value=%v\n",i,vVal.Field(i))
		tagVal := vTy.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Println("--->",tagVal)
		}
	}
	
	numM := vVal.NumMethod()
	fmt.Println(numM)
	vVal.Method(0).Call(nil)
	
	var para []reflect.Value
	para = append(para,reflect.ValueOf(45.6))
	vVal.Method(1).Call(para)
	para = append(para,reflect.ValueOf("韩梅梅"))
	vVal.Method(2).Call(para)
}

func main(){
	var stu Student = Student{Name:"张三",Age:12,No:"1111111",Score:0.0,}
	testReflect(stu)

}


























