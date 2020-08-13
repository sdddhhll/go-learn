package main

import "fmt"
import "reflect"

type Student struct{
	Name string
	Age int
}

func reflactTest01(b interface{}){
	a := reflect.TypeOf(b)
	fmt.Printf("%T\n",a)
	
	c := reflect.ValueOf(b)
	fmt.Printf("%T\n",c)
	
	d := c.Int()
	fmt.Printf("%T,%d\n",d,d)
	
	e := c.Interface()
	fmt.Printf("%T,%v\n",e,e)
	
	f := e.(int) //类型断言
	fmt.Printf("%T,%d\n",f,f)
	
}

func reflactTest02(b interface{}){
	fmt.Printf("%T,%v\n",b,b)
	c := reflect.ValueOf(b)
	fmt.Printf("%T,%v\n",c,c)
	d := reflect.TypeOf(b)
	fmt.Printf("%T,%v\n",d,d)
	e := b.(Student)
	fmt.Printf("%T,%v\n",e,e)
	f := e.Name
	fmt.Printf("%T,%v\n",f,f)
	
	g := c.Interface()
	h := g.(Student)
	fmt.Printf("%T,%v\n",h,h)
	
}

func test(){
	const a = 9
	fmt.Println(a)
	const b int = 9
	fmt.Println(b)
	const c int = 9/3
	fmt.Println(c)
	
	const (
		s = 9
		o = 9
	)
	fmt.Println(s)
	fmt.Println(o)
	const (
		x = iota  //初始值是0
		y //1
		z //2
	)
	
	fmt.Println(x)//0
	fmt.Println(y)//1
	fmt.Println(z)//2
	
	const (
		x1 = iota
		x2 = iota
		x3,x4 = iota,iota
		x5 = iota
		x6 = iota//iota作用的范围是在()内  如果新启一个()那么会从0开始
		x7
		x8
	)
	fmt.Println(x1,x2,x3,x4,x5,x6,x7,x8) //0 1 2 2 3 4 5 6
}


func main(){
	//reflactTest01(100)
	//stu := Student{Name:"tom",Age:90}
	//reflactTest02(stu)
	test()
}




















