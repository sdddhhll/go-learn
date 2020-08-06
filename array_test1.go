package main

import (
	"fmt"
)
func main(){
	var cha [26]byte 
	for i:=0;i<26;i++{
		cha[i]=(byte)(65+i)
	}
	for i:=0;i<26;i++{
		fmt.Printf("%c\n",cha[i])
	}
}
