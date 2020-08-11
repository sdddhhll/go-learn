package main

import "fmt"
import "reflect"

func reflactTest01(b interface{}){
	a := reflect.TypeOf(b)
	fmt.Printf("%T\n",a)
	
	c := reflect.ValueOf(b)
	fmt.Printf("%T\n",c)
	
	d := c.Int()
	fmt.Printf("%T,%d\n",d,d)
	
	e := c.Interface()
	fmt.Printf("%T,%v\n",e,e)
	
	f := e.(int)
	fmt.Printf("%T,%d\n",f,f)
	
}


func main(){
	reflactTest01(100)
	
}




















