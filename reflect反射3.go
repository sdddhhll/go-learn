package main
import "fmt"
import "reflect"


func testReflect(b interface{}){
	c := reflect.TypeOf(b)
	d := reflect.ValueOf(b)
	e := d.Interface()
	f := e.(float64)
	fmt.Printf("%T,%v\n",c,c)
	fmt.Printf("%T,%v\n",d,d)
	fmt.Printf("%T,%v\n",e,e)
	fmt.Printf("%T,%v\n",f,f)
}

func main(){
	//var a float64 = 1.2
	//testReflect(a)
	
	var ss string = "zs"
	b := reflect.ValueOf(&ss)
	b.Elem().SetString("tom")
	fmt.Println("--->",ss)
	
	
}


























