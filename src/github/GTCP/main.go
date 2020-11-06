package main

import "fmt" 
import "errors"

func Suma(a,b int) (int,error) {
	if a < b {
		return 0, errors.New("el primer valor es menor que el segundo valor a < b")
	}

		return a + b, nil
	
}


func main(){
	 fmt.Println("Hello World")
	 r,err := Suma(1,2)
	 if err != nil {
		 fmt.Println(err)
	 } else {
		fmt.Println(r)
	 }
} 