package main

import "fmt"

type close func(s string)

func invokeClose (c close){
    c("hello world")
}

func main(){
    f := func(s string){
        fmt.Print(s)

    }
	invokeClose(f)	
	//invokeClose(func(s string)){
	//	fmt.Println(s)
	//}
}